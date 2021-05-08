Mục đích:
- Start up.
- Nên cần: Backend tốt.
- Điều kiện đủ: Kinh nghiệm để code sao cho tốt để QC có thể test.
- Kinh nghiệm tổ chức project, quản trị dự án, communication.

Làm rõ:

- Start một cái biz về tech base on: Go Ecosystem.
- Cần skills set về:
  - Coding, đặc biệt về Go
    - Easy to Test
  - Team work, communication
  - Quản trị dự án meet deadline
  - Tech lead:
    - Biết cách chia các task về nhỏ.


Chốt lại:

- Coding
  - Simple
    - Composition
  - Complex
    - Maintain phuc tap
    - Giai thich cho nguoi khac no phuc tap
- Testing
  - Testable
  - Test Case
  - TDD (test driven development)
- DevOps
  - Automatic
- Communication
  - Slack
  - Calendar
  - Online meeting
  - Document
- Task management
  - Jira

# Nhu la mot phan mem tot (BE)

- Ổn định
  - Ít lỗi.
    - Không nên: em không reproduce được, nhờ chị thử lại.
    - Design đúng đắn.
      - Hiểu bản chất hoạt động của máy tính
      - Bản chất của network
        - HTTP Protocol
        - GRPC Protocal
        - MySQL
      - Bản chất của security
  - Khi có lỗi định danh, navigate được đến cái root cause của lỗi
    - Và fix được lỗi.
  - Performance tốt: response nhanh.
  - Throughput tốt (scale triệu khách hàng)
- Function
  - Sắp xếp
  - Gọi nó
  - Unit Test  (Unit = Function)
    - Testable
      - Predictable
        - Input có một output không đổi.
    - Khó.
      - Vì mình không có viết cái function `sum(a, b) =c`
- Software
  - Sắp xếp
  - Gọi nhau
  - Communication
    - Req/Res => Handle error.
    - Event
      - PubSub
      - Queue

# Nhu the nao la mot san pham tot (FE)

- Phu hop nhu cau nhu cau cua nguoi dung
- Co chi phi dau tu hop ly

- Sản phẩm tốt là sản phẩm có trải nghiệm người dùng tốt.
  - Họ không bực mình khi dùng sản phẩm của mình.
  - Thích.
  - Hiệu quả.

https://github.com/tpphu/golang-training/wiki

https://github.com/topics/go
https://github.com/avelino/awesome-go
https://github.com/Alikhll/golang-developer-roadmap
https://landscape.cncf.io/

UT cho cai UT
