package concurrent

/*
	Take a function name, number of go routines
	concurrent.Execute(func_input_handle,func_output_handle,thread_count)

	-- interface to function ✔️
		`reflect.ValueOf(input_handler).Call([]reflect.Value{"string"})`
	-- interface to slice/channel [working]
	S
*/

import (
	"reflect"
	"sync"
)

func Execute(input_handler_interface interface{}, input_set_interface interface{}, thread_count int) {
	var wg sync.WaitGroup

	input_set := reflect.ValueOf(input_set_interface)         // interface to reflect.Value
	input_handler := reflect.ValueOf(input_handler_interface) // interface to reflect.Value

	ok := true
	var elem reflect.Value

	// typical way of implementing threading in golang
	for i := 0; i < thread_count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			/* this code loops through the channel and send
			   the data as input to `input_handler` function */
			for ok {
				if elem, ok = input_set.Recv(); ok {
					input_handler.Call([]reflect.Value{elem})
				}
			}
		}()
	}
	wg.Wait()
}
