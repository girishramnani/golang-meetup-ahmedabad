
// START1 OMIT
func Stream(parent context.Context, out chan<- Value) error {
  cctx, _ := context.WithDeadline(parent, time.Now().Add(time.Second * 5)) // HL
  for {
    v, err := DoSomething(ctx)
    if err != nil {
      return err
    }
    select {
      case <- cctx.Done(): // HL
      //check if context is canceled
        return cctx.Err()
      case out <- v:
    }
  }
}
// STOP1 OMIT

// START2 OMIT
func Stream(parent context.Context, out chan<- Value) error {
  cctx, cancel := context.WithCancel(parent) // HL

  go func() { // HL
    timer := time.AfterFunc(time.Second * 5, cancel) // HL
  }() // HL

  for {
    v, err := DoSomething(ctx)
    if err != nil {
      return err
    }
    select {
      case <- cctx.Done(): // HL
      //check if context is canceled
        return cctx.Err()
      case out <- v:
    }
  }
}
// STOP2 OMIT

// START3 OMIT
func Stream(ctx context.Context, out chan<- Value) error {
  for {
    v, err := DoSomething(ctx) // HL
    if err != nil {
      return err
    }
    select { // HL
      case <- ctx.Done(): // HL
        return ctx.Err() // HL
      case out <- v: // HL
    } //HL
  }
}
...
..
.
[Caller function]
ctx := context.Background()
go Stream(ctx, outChan)
// STOP3 OMIT

// START4 OMIT
func Stream(out chan<- Value) error {
    for {
      v, err := DoSomething() // HL
      if err != nil {
        return err
      }
      out <- v // HL
    }
  }
// STOP4 OMIT

// START5 OMIT
  type middlewareTypes int
  RequestIdType middlewareTypes = 1

  func decorator(f http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
      rnd := rand.Int63()
      ctx := r.Context() // HL
      newCtx := context.WithValue(ctx, RequestIdType, rnd) // HL
      f(w, r.WithContext(newCtx)) // HL
    }
  }

  //somewhere inside a handler
  ctx := r.Context() // HL
  v, ok := ctx.Value(RequestIdType).(int); // HL
// STOP5 OMIT

// START6 OMIT
func DoSum() float64
{
  db, _ := sql.Open("oci8", "connection string...")
  ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*500)) // HL
  defer cancel() // HL
  ch = make(chan float64)
  go makeDbCall(ctx, db, ch)
  select {
    case <-ctx.Done():
        fmt.Print(ctx.Err())
        return
    case sum := <- ch:
        return sum
  }
}
func makeDbCall(ctx context.Context, db *sql.DB, out chan<- float64) {
  ok, deadline := ctx.Deadline()
  //check if deadline is near and decide not to proceed
  row := db.QueryRowContext(ctx, "select sum(balance) from accounts") // HL
  var sum float64
  err := row.Scan(&sum)
  checkErr(err)
  out <- sum // HL
}
// STOP6 OMIT
