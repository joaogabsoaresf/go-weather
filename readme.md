# Projeto Go com Redis

Este é um projeto em GO para consultar o clima de uma cidade por uma API simples, o projeto utiliza o Redis para salvar em cache os dados em memória. O objeto salvo no Redis vai existir por 10 minutos. A consulta é feita na API do OpenWeather

## Clonando o Projeto

Para clonar este projeto para o seu ambiente local, execute o seguinte comando:

```bash
git clone https://github.com/joaogabsoaresf/go-weather
```
Instalando Dependências
Certifique-se de ter o Go instalado em sua máquina. Em seguida, você precisará instalar as dependências do projeto. Para isso, use o seguinte comando:

```bash
go mod tidy
```
Isso instalará todas as dependências listadas no arquivo go.mod.

Certifique-se também de ter todas as variaveis de ambiente necessárias para o funcionamento do projeto.
Tenha a URL de conexão do Redis e a chave de API do OpenWeather em mãos, digite o comando, e cole no destino:

```bash
cp .env.example .env
```

Executando o Projeto
Depois de clonar o projeto e instalar as dependências, você pode executá-lo localmente. Navegue até o diretório do projeto e execute o seguinte comando:

```bash
go run main.go
```
Isso iniciará o projeto e você verá os resultados no console.

## Endpoints
O projeto no momento só possui um endpoint e ele estará disponivel em:
http://localhost:8080/api/v1/weather/NOMEDACIDADE

Basta alterar o nome da cidade e buscar, o valor será salvo no Redis para retornar mais rápido nas próximas consultas.
