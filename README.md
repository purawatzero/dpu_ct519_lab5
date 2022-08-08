วิธีการ Deploy 
1. Download Package และ extract ลงบน path ที่ต้องการ
2. Run command เพื่อ Deploy golang container : 
    sudo docker compose up -d
3. Run Start service golang เพื่อ start webservice
    sudo docker exec -i ctn_golang sh -c 'go run main.go'
4. เปิด browser เพื่อทดสอบ webserver 
    http://localhost:8080
