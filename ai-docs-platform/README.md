# AI Software Documentation Platform 🚀

Nền tảng AI phân tích mã nguồn và tự động tạo tài liệu kỹ thuật hoạt động hoàn toàn bằng mô hình AI cục bộ (Local AI).

## 🌟 Trạng thái Dự án (Project Status)
Dự án đã hoàn thành **Phase 1** và đang thực hiện **Phase 2 & 3** theo đúng lộ trình WBS. Nền tảng hiện đã có khung hạ tầng vững chắc với Go Backend và Next.js Frontend.

---

## ✅ Những gì đã hoàn thành (Achieved)

### 1. Kiến trúc & Hạ tầng (Infra)
- [x] Khởi tạo cấu trúc dự án chuẩn (Go Backend + Next.js Frontend).
- [x] Thiết lập **Docker Compose** cho toàn bộ stack: PostgreSQL, Qdrant (Vector DB), Redis.
- [x] Cấu hình tích hợp **Ollama** cho việc chạy LLM cục bộ (Local LLM).

### 2. Backend Core (Go)
- [x] Xây dựng Server hiệu năng cao bằng **Gin Gonic**.
- [x] Triển khai **File Ingestion Pipeline**:
    - Nhận file ZIP mã nguồn qua API.
    - Tự động giải nén và quét cấu trúc thư mục (Sandbox).
    - Nhận diện các Framework (Node.js, Go, Java).
- [x] Triển khai **RAG Pipeline Core**:
    - Logic **Chunking** mã nguồn thông minh.
    - Tích hợp **Embedding** cục bộ qua Ollama (`nomic-embed-text`).
    - Lưu trữ Vector và Metadata vào **Qdrant** sử dụng gRPC để đạt tốc độ tối đa.
- [x] Xây dựng khung **Multi-Agent AI**:
    - Khởi tạo Base Agent và specialized agent đầu tiên (**Code Analyzer**).

### 3. Frontend Dashboard (Next.js)
- [x] Thiết kế giao diện **Premium Dark Mode** với phong cách hiện đại.
- [x] Xây dựng Dashboard quản lý dự án, Project Cards và Navbar.
- [x] Tích hợp hệ thống Icon chuyên nghiệp (Lucide).

---

## 🛠️ Công nghệ sử dụng
- **Frontend**: Next.js 14, TailwindCSS, Lucide Icons.
- **Backend**: **Go (Golang)**, Gin Gonic, gRPC.
- **AI/ML**: Ollama, Qdrant (Vector Database).
- **Storage**: PostgreSQL, Redis.

---

## 🚧 Kế hoạch tiếp theo (Upcoming Tasks)
- [ ] Triển khai **RAG Retrieval**: Tìm kiếm đoạn code liên quan nhất dựa trên ngữ cảnh tài liệu.
- [ ] Hoàn thiện các AI Agents chuyên biệt: **API Mapper**, **DB Reader**, **Content Writer**.
- [ ] Xây dựng bộ **Prompt Templates** cho 16 loại tài liệu đầu ra (README, SRS, Báo cáo thực tập...).
- [ ] Tích hợp **Rich Text Editor** trên Frontend để chỉnh sửa tài liệu trực tiếp.
- [ ] Phát triển công cụ xuất file **DOCX/PDF**.

---

## 🚀 Hướng dẫn khởi động nhanh

1. **Chuẩn bị Ollama**:
   ```bash
   ollama pull qwen2.5-coder:7b
   ollama pull nomic-embed-text
   ```

2. **Chạy hạ tầng**:
   ```bash
   docker-compose up -d
   ```

3. **Chạy Backend (Go)**:
   ```bash
   cd go-backend
   go run .
   ```

4. **Chạy Frontend (Next.js)**:
   ```bash
   cd frontend
   npm run dev
   ```

---
*Cập nhật lần cuối: 15/05/2026*
