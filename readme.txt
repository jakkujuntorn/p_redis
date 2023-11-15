
-สถานตอนนี้
    -Func New ที่ REturn struct จะ Return productRepositoryDB หรือ &productRepositoryDB ก็ได้
        -productRepositoryDB จะ return interface *******
        -&productRepositoryDB จะ return struct *********

- handler
    -layer นี้จะรับ gin หรือ fiber
    -layer นี้จะไม่ return เพราะมันจะออกไปทาง response json 

- Func New พารามิเตอร์ การ return ต้องเรียงว่าอะไรมาก่อนมาหลังด้วย ******

-การใชเงาน Redis
    -จะอยู่ที่ services  *****
    
************************************************************************
-docker 
    - docker compose run เพื่อ run โปรแกรมตัวอื่นใน containner
    -docker compose run --rm  / run แล้วทำลายทิ้ง โหลดเทสเสร็จ ืำลายทิ้ง
        -docker compose run --rm k6 run 
    - run บางตัว docker compose up ตามด้วนโปรแรกมที่จะรัน influxdb grafana 
    - ปิด dockerใช้คำสั่ง docker compose down ******
-redis
    - docker exec -it *ชื่อ conainner* sh 
    - docker compose up redis 
    - redis ควร ว่างที่ services มากกว่า repository *****
    - สร้างที่เก็บข้อมูลไว้ที่ data/redis/appendonly  ต้องไปเซตใน redis.conf ****
-mariadb
    -ใช้ ports: 3366 / ใช้ 3306 error ******

-k6 (Load Test  ต้อง run go ก่อน)
    -run นอก docker และให้เก็บใน influxdb  แต่ถ้า user เยอะๆ จะ error
        - k6 run ./scripts/test.js -o influxdb=http://localhost:8086*k6
    -run ผ่าน docker
        - ที่อยู่ scripts ที่จะ รัน คือ /scripts/test.js ****
        - docker compose run --rm k6 run /scripts/test.js -u5 -d5s (u คือ จำนวน user, d คือ ระยะเวลา -d1m5s 1นาที5วิ ก็ได้)
        - docker compose run --rm k6 run /scripts/test.js  (เซต option ไว้แล้วเลยไม่ต้อง -u -d)

 -influxdb
    -อยู๋ใน docker (k6 ต้องอยู่ใน docker ด้วย เวลายิงข้อมูลเยอะๆจะสะดวกกว่า)

-gafana
    -อยู่ใน docker พร้อมใช้



****  ต้องแยกกัน run   run พร้อมกันไม่ทำงาน **********

-สร้าง MockDatabase ด้วย gorm ได้