# AI Software Documentation Platform 🚀

Nền tảng AI phân tích mã nguồn và tự động tạo tài liệu kỹ thuật hoạt động hoàn toàn bằng mô hình AI cục bộ (Local AI).

## 🌟 Tính năng chính
- **Local AI First**: Không phụ thuộc vào API bên ngoài (OpenAI, Claude). Sử dụng Ollama.
- **Phân tích đa nền tảng**: Hỗ trợ Next.js, React, Spring Boot, Flask, Express, Go, v.v.
- **RAG Engine**: Truy xuất thông tin thông minh từ codebase bằng Qdrant.
- **Multi-Agent Architecture**: Các agent chuyên biệt cho Code, DB, API và Writing.
- **Xuất tài liệu**: DOCX, PDF, Markdown cho 16 loại tài liệu kỹ thuật.

## 🛠️ Công nghệ sử dụng
- **Frontend**: Next.js 14, TailwindCSS, Lucide Icons.
- **Backend**: Go (Gin Gonic), LangChain/CrewAI logic (via Go agents).
- **Vector DB**: Qdrant.
- **Database**: PostgreSQL.
- **Inference**: Ollama (Qwen2.5-Coder, DeepSeek).

## 🚀 Hướng dẫn khởi động

### 1. Chuẩn bị môi trường
- Cài đặt [Docker Desktop](https://www.docker.com/products/docker-desktop/).
- Cài đặt [Ollama](https://ollama.com/) và tải model:
  ```bash
  ollama pull qwen2.5-coder:7b
  ollama pull nomic-embed-text
  ```

### 2. Khởi động hạ tầng (Database, Vector DB)
```bash
docker-compose up -d
```

### 3. Chạy Backend (Go)
```bash
cd go-backend
go run .
```

### 4. Chạy Frontend
```bash
cd frontend
npm install
npm run dev
```

Truy cập dashboard tại: `http://localhost:3000`

---
*Dự án được xây dựng nhằm hỗ trợ lập trình viên và doanh nghiệp tự động hóa quy trình viết tài liệu mà vẫn đảm bảo quyền riêng tư dữ liệu.*
