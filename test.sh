echo "-----------------day 2-----------------"
echo "test GET /"
curl -i http://localhost:9999/
echo "test GET /hello"
curl "http://localhost:9999/hello?name=attackoncs"
echo "test POST /login"
curl "http://localhost:9999/login" -X POST -d 'username=attackoncs&password=1234'
echo "test /Not"
curl "http://localhost:9999/Not"
echo "-----------------day 2-----------------"
echo "-----------------day 3-----------------"
curl "http://localhost:9999/hello/attackoncs"
curl "http://localhost:9999/assets/css/attackoncs.css"
echo "-----------------day 3-----------------"