package main
import(
	"testing"
	"bytes"
	"net/http"
	"net/http/httptest"
	
)
func TestAddemp(t *testing.T){
	var jsonstr =[]byte(`{"name":"shry","department":"cs","address":"bang","skills":["golang","java"]}`)
	req, err := http.NewRequest("POST", "/add", bytes.NewBuffer(jsonstr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Addemp)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	


}
func TestDeleteemp(t *testing.T){
	//please execute both cases related to delete conditions

	//var jsonstr =[]byte(`{"empid":31,"permanentlyDelete":true}`)//employee has to be permanently deleted

	var jsonstr =[]byte(`{"empid":4}`)//employee has to be deactivated

	req, err := http.NewRequest("POST", "/delete", bytes.NewBuffer(jsonstr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Deleteemp)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	


}
func TestListemp(t *testing.T){
		var jsonstr =[]byte(`{"empid": 5}`)//employee has to be deactivated

	req, err := http.NewRequest("POST", "/list", bytes.NewBuffer(jsonstr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Listemp)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	
// expected := `[{"empid": 5,"name": "","department": "cs","address": "bang","skills": ["golang",""]}]`
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}

}
func TestRestoreemp(t *testing.T){
	var jsonstr =[]byte(`{"empid": 5}`)//employee has to be deactivated

req, err := http.NewRequest("POST", "/restore", bytes.NewBuffer(jsonstr))
if err != nil {
	t.Fatal(err)
}
req.Header.Set("Content-Type", "application/json")
rr := httptest.NewRecorder()
handler := http.HandlerFunc(Restoreemp)
handler.ServeHTTP(rr, req)
if status := rr.Code; status != http.StatusOK {
	t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
}



}
///////////////////////////////////////
func TestSearchemp(t *testing.T){
	var jsonstr =[]byte(`{"empid": 5}`)//employee has to be deactivated

req, err := http.NewRequest("POST", "/search", bytes.NewBuffer(jsonstr))
if err != nil {
	t.Fatal(err)
}
req.Header.Set("Content-Type", "application/json")
rr := httptest.NewRecorder()
handler := http.HandlerFunc(Searchemp)
handler.ServeHTTP(rr, req)
if status := rr.Code; status != http.StatusOK {
	t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
}

 //expected := `[{"empid": 5,"name": "","department": "cs","address": "bang","skills": ["golang",""]}]`
 
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}

}
////////////////////////////////////
func TestUpdateemp(t *testing.T){
	var jsonstr =[]byte(`{"empid":20,"department":"cs","address":"bang","skills":["golang","java","c"]}`)//employee has to be deactivated

req, err := http.NewRequest("POST", "/update", bytes.NewBuffer(jsonstr))
if err != nil {
	t.Fatal(err)
}
req.Header.Set("Content-Type", "application/json")
rr := httptest.NewRecorder()
handler := http.HandlerFunc(Updateemp)
handler.ServeHTTP(rr, req)
if status := rr.Code; status != http.StatusOK {
	t.Errorf("handler returned wrong status code: got %v want %v",
		status, http.StatusOK)
}



}