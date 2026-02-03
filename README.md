𝟏. 𝐋ớ𝐩 𝐃𝐨𝐦𝐚𝐢𝐧 (𝐋õ𝐢 𝐧𝐠𝐡𝐢ệ𝐩 𝐯𝐮̣)
Đây là phần lõi của hệ thống, 100% C# thuần túy.
📁 Domain
├── 📁 DomainEvents      # Sự kiện Domain
├── 📁 Entities          # Thực thể
├── 📁 Enumerators       # Liệt kê
├── 📁 Constants         # Hằng số
├── 📁 Exceptions        # Ngoại lệ
├── 📁 Repositories      # Giao diện Repository
├── 📁 Shared           # Thành phần dùng chung
└── 📁 ValueObjects     # Đối tượng Giá trị
𝟐. 𝐋ớ𝐩 Ứ𝐧𝐠 𝐝𝐮̣𝐧𝐠
Lớp này định nghĩa hệ thống của bạn nên làm gì.
📁 Application
├── 📁 Abstractions     # Giao diện trừu tượng
│  ├── 📁 Data
│  ├── 📁 Email
│  └── 📁 Messaging
├── 📁 Behaviors        # Pipeline Behaviors (MediatR)
├── 📁 Contracts        # DTOs, Requests, Responses
├── 📁 User             # Ngữ cảnh Người dùng
│  ├── 📁 Commands
│  └── 📁 Queries
├── 📁 Order            # Ngữ cảnh Đơn hàng
│  ├── 📁 Commands
│  └── 📁 Queries
└── 📁 UseCases         # (Tùy chọn) Các trường hợp sử dụng
𝟑. 𝐋ớ𝐩 𝐂ơ 𝐬ở 𝐡ạ 𝐭ầ𝐧𝐠
Đây là nơi chứa các chi tiết công nghệ. Nó triển khai tất cả các giao diện trừu tượng.
📁 Infrastructure
├── 📁 Data                      # Dữ liệu
│  ├── 📁 Repositories          # Triển khai Repository
│  ├── 📁 Migrations            # Migration Database
│  └── 📁 DataContext
│    └── ApplicationDbContext.cs
├── 📁 Messaging                # Triển khai Message Queue, v.v.
├── 📁 Services                 # Các dịch vụ cụ thể (Email, File, v.v.)
└── 📁 Jobs                     # Công việc nền (Background Jobs)
𝟒. 𝐋ớ𝐩 𝐓𝐫ì𝐧𝐡 𝐛à𝐲
Đây là điểm đầu vào: Controllers, Endpoints.
📁 Presentation
├── 📁 Controllers      # Controller (API/MVC)
├── 📁 Middlewares      # Middleware
├── 📁 Extensions       # Các phương thức mở rộng
├── 📁 Endpoints        # Dành cho Minimal APIs (tùy chọn)
└── 📁 ViewModels       # Dành cho frontend hoặc UI (tùy chọn)
