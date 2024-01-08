# Use a imagem oficial do pgAdmin4
FROM dpage/pgadmin4

# Configurações padrão do pgAdmin4
ENV PGADMIN_DEFAULT_EMAIL=titiura@gmail.com
ENV PGADMIN_DEFAULT_PASSWORD=47523146phB@

# Exponha a porta 80
EXPOSE 80

# Crie diretórios de volumes para persistir dados
VOLUME [ "/var/lib/pgadmin", "/var/log/pgadmin" ]

# Comando de inicialização do contêiner
CMD ["./entrypoint.sh"]