//package funding
//
//import "testing"
//
//func BenchmarkFund(b *testing.B) {
//    // Add as many dollars as we have iterations this run
//    fund := NewFund(b.N)
//
//    // Burn through them one at a time until they are all gone
//    for i := 0; i < b.N; i++ {
//        fund.Withdraw(1)
//    }
//
//    if fund.Balance() != 0 {
//        b.Error("Balance wasn't zero:", fund.Balance())
//    }
//}
package funding

import (
    "sync"
    "testing"
)

const WORKERS = 10

func BenchmarkWithdrawals( b *testing.B) {
    //skip N = 1

    if b.N < WORKERS {
        return
    }

    //Add as many mbesha as we have iterations this run
    fund := NewFund(b.N)

    //assume it divides cleanly
    dollarsPerFounder := b.N / WORKERS

    var wg sync.WaitGroup

    for i :=0;i<WORKERS; i++ {
        //notify the waitgroup were adding a go routine
        wg.Add(1)
        //spawn off a founder worker, as a closure
        go func() {

            defer wg.Done()

            for i := 0; i < dollarsPerFounder; i++ {
                fund.Withdraw(1)
            }
        }()
    }

    wg.Wait() //wait for all workers to finish

    if fund.Balance() != 0 {
        b.Error("Balance wasn't zero:", fund.Balance())
    }


}