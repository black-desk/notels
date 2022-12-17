package conn_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/black-desk/notels/pkg/jsonrpc2"
	. "github.com/black-desk/notels/pkg/jsonrpc2/conn"
	"github.com/black-desk/notels/pkg/jsonrpc2/message"
	"github.com/black-desk/notels/pkg/jsonrpc2/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sourcegraph/conc/pool"
)

func TestConn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Conn Suite")
}

var _ = Describe("JSON-RPC 2.0 Conn", func() {
	var err error

	Context("created", func() {
		Context("without transfer", func() {
			It("will failed", func() {
				_, err = New()
				Expect(err).NotTo(BeNil())
				// TODO(black_desk): errors.Is and errors.As is not working with
				// validator, just skip this check.
			})
		})
		Context("with mock transfer", func() {
			var (
				mockCtrl     *gomock.Controller
				mockTransfer *mocks.MockTransfer
			)

			BeforeEach(func() {
				By("Create a mock transfer", func() {
					mockCtrl = gomock.NewController(GinkgoT())
					mockTransfer = mocks.NewMockTransfer(mockCtrl)
				})
			})

			AfterEach(func() {
				mockCtrl.Finish()
			})

			It("should success", func() {
				_, err = New(WithTransfer(mockTransfer))
				Expect(err).To(BeNil())
			})

			Context("and context", func() {
				var (
					mockExpect *mocks.MockTransferMockRecorder
					conn       *Conn
				)

				BeforeEach(func() {
					mockExpect = mockTransfer.EXPECT()

					conn, err = New(
						WithTransfer(mockTransfer),
						WithContext(context.Background()),
					)
				})

				AfterEach(func() {
					mockCtrl.Finish()
				})

				It("should success", func() {
					Expect(err).To(BeNil())
				})

				When("no adaptor registered", func() {
					Context("and Conn.Loop called", func() {
						var (
							reqChan  chan json.RawMessage
							p        *pool.ErrorPool
							done     chan struct{}
							rawResps []json.RawMessage
						)

						BeforeEach(func() {
							By("Mock Transfer.Run")
							p = pool.New().WithErrors()
							done = make(chan struct{})

							runCalled := mockExpect.Run().DoAndReturn(func() error {
								<-done
								return ErrCanceled
							})

							By("Mock Transfer.Read")
							reqChan = make(chan json.RawMessage)
							mockExpect.Read().DoAndReturn(
								func() <-chan json.RawMessage {
									return reqChan
								}).AnyTimes().After(runCalled)

							p.Go(conn.Loop)
						})

						AfterEach(func() {
							By("Mock transfer", func() {
								mockExpect.Close().DoAndReturn(func() {
									close(done)
									close(reqChan)
									return
								})
							})
							By("Conn.Cancel called", func() {
								err = conn.Cancel()
								Expect(err).To(BeNil())
							})
							By("Wait Conn.Loop return and response collecting end", func() {
								err = p.Wait()
								By("There should be an ErrCanceled returned by Loop", func() {
									Expect(err).To(MatchError(ErrCanceled))
								})
								By("And no response sended", func() {
									Expect(rawResps).To(BeEmpty())
								})
							})
						})

						Context("then an invalid message", func() {
							Context(`not contain "id" (for example`, func() {
								messages := []string{"{}", `""`}
								for i := range messages {
									Context(fmt.Sprintf("%s) arrives", messages[i]), func() {
										i := i
										BeforeEach(NodeTimeout(time.Second), func(ctx context.Context) {
											select {
											case reqChan <- []byte(messages[i]):
											case <-ctx.Done():
												Fail("Timeout")
											}
										})
										It("should have no response sended", func() {})
									})
								}
							})
							Context(`contains "id" (for example`, func() {
								messages := []string{`{"id":123}`, `{"id":"456"}`}
								for i := range messages {
									Context(fmt.Sprintf("%s) arrives", messages[i]), func() {
										i := i
										var respChan chan *jsonrpc2.MessageToWrite
										BeforeEach(NodeTimeout(time.Second), func(ctx context.Context) {
											By("Mock Transfer.Write", func() {
												respChan = make(chan *jsonrpc2.MessageToWrite)
												mockExpect.Write().DoAndReturn(
													func() chan<- *jsonrpc2.MessageToWrite {
														return respChan
													}).AnyTimes()
											})
											select {
											case reqChan <- []byte(messages[i]):
											case <-ctx.Done():
											}
										})

										It("should send an error as response", func(ctx context.Context) {
											var m *jsonrpc2.MessageToWrite
											select {
											case m = <-respChan:
											case <-ctx.Done():
												Fail("Timeout")
											}

											msg := struct{ Error struct{ Code int } }{}
											err = json.Unmarshal(m.Msg, &msg)

											Expect(err).To(BeNil())
											Expect(msg.Error.Code).To(Equal(message.CodeInvalidRequest))

											close(respChan)
										})
									})
								}
							})
						})

					})
				})

			})
		})
	})
})
