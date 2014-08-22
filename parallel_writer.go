package main

type ParallelWriter struct {
	writers []Writer
}

func (writer *ParallelWriter) Init() error {

	for _, w := range writer.writers {
		err := w.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

func (writer *ParallelWriter) Close() error {

	for _, w := range writer.writers {
		err := w.Close()
		if err != nil {
			return err
		}		
	}
	return nil
}

func (writer *ParallelWriter) Write(b []byte) (n int, err error) {
	
	n = 0
	for _, w := range writer.writers {
		num, err := w.Write(b)
		n += num
		if err != nil {
			return n, err
		}
	}

	return n, nil
}
