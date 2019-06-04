package main

var (
	testTempDir   = "temp_test"
	testTempBuild = "testbuild"
)

// func Test_muxWithHandler(t *testing.T) {

// 	testSuite := []struct {
// 		flag string
// 	}{
// 		{
// 			"yaml",
// 		},
// 	}

// 	for _, test := range testSuite {

// 		t.Run(fmt.Sprintf("Testing muxHanlder with %v flag", test.flag), func(t *testing.T) {
// 			_, err := muxWithHandler("yaml")
// 			if err != nil {
// 				t.Errorf("Wanted error to be nil, got %v", err)
// 			}
// 		})
// 	}

// 	t.Run("Testing muxHanlder with invalid parser type", func(t *testing.T) {
// 		_, err := muxWithHandler("nonExistingParserType")

// 		if err == nil {
// 			t.Errorf("Wanted error not to be nil")
// 		}
// 	})

// 	t.Run("Testing fallback handler gets called", func(t *testing.T) {
// 		mux, _ := muxWithHandler("yaml")
// 		srv := httptest.NewServer(mux)

// 		res, err := http.Get(srv.URL + "/nonexistingroute")
// 		if err != nil {
// 			t.Errorf("Error performing http req on test server : %v", err)
// 		}

// 		body, err := ioutil.ReadAll(res.Body)
// 		if err != nil {
// 			t.Errorf("Could not read body : %v", err)
// 		}

// 		if string(body) != string(defaultResponse) {
// 			t.Errorf("Wanted %v as a default response, got %v", defaultResponse, string(body))
// 		}

// 	})

// }
