echo "Port: 5432"
echo "dbNam : postgres" 
echo "user : shail"
docker run -itd -e POSTGRES_USER=shail -e POSTGRES_PASSWORD=shail -p 5432:5432 -v data:/var/lib/postgresql/data --name postgresql postgres