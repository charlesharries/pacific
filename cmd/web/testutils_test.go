package main

// func newTestApplication(t *testing.T) *application {
// 	templateCache, err := newTemplateCache("./../../resources/views")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	session := sessions.New([]byte("abc123"))
// 	session.Lifetime = 12 * time.Hour

// 	return &application{
// 		errorLog:      log.New(ioutil.Discard, "", 0),
// 		infoLog:       log.New(ioutil.Discard, "", 0),
// 		session:       session,
// 		templateCache: templateCache,
// 		users:         &mock.UserModel{},
// 	}
// }
