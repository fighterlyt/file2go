package compress

import (
	"bytes"
)
const (
	x855595167=`HttpURLConnection httpConn = (HttpURLConnection) myURL.openConnection();
httpConn.setDoOutput(true);
httpConn.setDoInput(true);
httpConn.setUseCaches(false);
httpConn.setRequestMethod("POST");
httpConn.setRequestProperty("Content-Type", "application/json;charset=utf-8");
httpConn.setConnectTimeout(5000);
httpConn.setReadTimeout(2 * 5000);
	`
)
type File2GoModel struct{
	buffer *bytes.Buffer
}
func NewFile2GoModel() File2GoModel{
	return File2GoModel{
		buffer:bytes.NewBuffer([]byte(x855595167)),
	}
}
func(m File2GoModel)Read(p []byte) (n int, err error){
	return m.buffer.Read(p)
}
