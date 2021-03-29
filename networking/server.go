package networking
//
//import (
//	"log"
//	"net"
//	"sync"
//)
//
//type Server struct {
//	listener net.Listener
//	quit     chan interface{}
//	wg       sync.WaitGroup
//}
//
//
//func NewServer(addr string) *Server {
//	s := &Server{
//		quit: make(chan interface{}),
//	}
//	l, err := net.Listen("tcp", addr)
//	if err != nil {
//		log.Fatal(err)
//	}
//	s.listener = l
//	s.wg.Add(1)
//	s.serve()
//	return s
//}
//
//func (s *Server) serve() {
//	defer s.wg.Done()
//
//	for {
//		conn, err := s.listener.Accept()
//		if err != nil {
//			select {
//			case <-s.quit:
//				return
//			default:
//				log.Println("accept error", err)
//			}
//		} else {
//			s.wg.Add(1)
//			go func() {
//				s.handleConnection(conn)
//				s.wg.Done()
//			}()
//		}
//	}
//}
//
//func (s *Server) Stop() {
//	close(s.quit)
//	s.listener.Close()
//	s.wg.Wait()
//}
