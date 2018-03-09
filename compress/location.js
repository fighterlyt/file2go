HttpURLConnection httpConn = (HttpURLConnection) myURL.openConnection();
httpConn.setDoOutput(true);
httpConn.setDoInput(true);
httpConn.setUseCaches(false);
httpConn.setRequestMethod("POST");
httpConn.setRequestProperty("Content-Type", "application/json;charset=utf-8");
httpConn.setConnectTimeout(5000);
httpConn.setReadTimeout(2 * 5000);
	