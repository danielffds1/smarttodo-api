# SmartTodo+ API

API REST inteligente de gerenciamento de tarefas com integração de dados climáticos e sugestões baseadas em IA

## Sobre o Projeto

SmartTodo+ é uma API REST desenvolvida em Go que combina gerenciamento de tarefas com inteligência artificial e dados climáticos em tempo real. O sistema analisa suas atividades planejadas e fornece sugestões personalizadas baseadas nas condições meteorológicas da sua cidade.

## ✨ Features Principais

🔐 Autenticação JWT completa com refresh tokens  
📝 CRUD de Tarefas com filtros avançados e paginação  
🌤️ Integração OpenWeather para dados climáticos em tempo real  
🤖 IA Generativa (Gemini) para sugestões contextuais  
🎯 Clean Architecture bem estruturada  
🔍 Busca textual e filtros dinâmicos  
📊 Paginação com metadados  
🗄️ PostgreSQL com GORM  
🐳 Docker Compose para ambiente de desenvolvimento

## 🛠️ Tecnologias Utilizadas

### Backend
- **Go 1.21+** - Linguagem principal
- **Chi Router** - Framework HTTP minimalista e performático
- **GORM** - ORM para Go
- **PostgreSQL 15** - Banco de dados relacional
- **JWT** - Autenticação stateless
- **Bcrypt** - Hash de senhas

### Integrações
- **OpenWeather API** - Dados climáticos em tempo real
- **Gemini AI (Google)** - Geração de sugestões inteligentes

### DevOps
- **Docker & Docker Compose** - Containerização
- **Air** - Live reload para desenvolvimento
- **Make** - Automação de comandos

## 🏗️ Arquitetura

Este projeto segue os princípios de Clean Architecture com separação clara de responsabilidades:

smarttodo-api/
├── cmd/api/ # Entry point da aplicação
├── internal/
│ ├── domain/ # Entidades e regras de negócio
│ ├── usecase/ # Casos de uso
│ └── infrastructure/ # Implementações (DB, HTTP, APIs externas)
└── pkg/ # Código reutilizável