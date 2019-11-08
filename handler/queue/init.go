package queue

func init(){
	go func() {
		weather()
	}()
}