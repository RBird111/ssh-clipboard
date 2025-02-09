package clipboard

import (
	"log"
	"net"
	"sync"
)

type ServerOpts struct {
	Address string
	ClipCmd
}

func NewServer(opts ServerOpts) *ClipServer {
	l, err := net.Listen("tcp", opts.Address)
	if err != nil {
		log.Fatal(err)
	}
	s := &ClipServer{
		listener: l,
		quit:     make(chan any),
		clipCmd:  opts.ClipCmd,
	}
	s.wg.Add(1)
	go s.serve()
	return s
}

type ClipServer struct {
	listener net.Listener
	quit     chan any
	wg       sync.WaitGroup
	clipCmd  ClipCmd
}

func (s *ClipServer) Stop() {
	close(s.quit)
	s.listener.Close()
	s.wg.Wait()
}

func (s *ClipServer) serve() {
	defer s.wg.Done()
	log.Println("clipboard server started")

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			select {
			case <-s.quit:
				log.Println("clipboard server shutting down...")
			default:
				log.Println("accept error:", err)
			}
			return
		}

		s.wg.Add(1)
		go s.handleConnection(conn)
	}
}

func (s *ClipServer) handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		s.wg.Done()
	}()

	clip, err := NewClipboard(s.clipCmd)
	if err != nil {
		log.Println("clipboard error:", err)
		return
	}

	if err := clip.CopyFrom(conn); err != nil {
		log.Println("clipboard copy error:", err)
	}
}
