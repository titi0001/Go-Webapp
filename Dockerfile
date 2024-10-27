# Use a imagem oficial do pgAdmin4 como base
FROM dpage/pgadmin4

# Exponha a porta 3000 para acessar o pgAdmin4
EXPOSE 3000

# Configure o pgAdmin4 para iniciar na porta 3000
ENV PGADMIN_LISTEN_PORT=3000

# Configure a URL de conexão para o PostgreSQL
# Aqui você pode adicionar a lógica para fornecer as credenciais de conexão ao PostgreSQL
# Este exemplo usa variáveis de ambiente para simplificar, mas é recomendável usar um método mais seguro para fornecer credenciais
ENV PGADMIN_DEFAULT_EMAIL=admin@pgadmin.org
ENV PGADMIN_DEFAULT_PASSWORD=admin

# Comando para iniciar o pgAdmin4 quando o contêiner for iniciado
CMD [ "python", "/usr/local/lib/python3.7/site-packages/pgadmin4/pgAdmin4.py" ]