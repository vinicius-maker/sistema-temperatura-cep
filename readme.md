# Desafio - Sistema de temperatura por CEP

## Objetivo
Desenvolver um sistema em Go que receba um CEP, identifique a cidade e retorne o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin). Este sistema deverá ser publicado no Google Cloud Run.

## Requisitos

- O sistema deve receber um CEP válido de 8 dígitos.
- O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin.
- O sistema deve responder adequadamente nos seguintes cenários:
    - **Em caso de sucesso:**
        - Código HTTP: `200`
        - Response Body:
          ```json
          { "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
          ```
    - **Em caso de falha (CEP inválido):**
        - Código HTTP: `422`
        - Mensagem: `invalid zipcode`
    - **Em caso de falha (CEP não encontrado):**
        - Código HTTP: `404`
        - Mensagem: `can not find zipcode`
- Deverá ser realizado o deploy no Google Cloud Run.

## Dicas
- Utilize a API [viaCEP](https://viacep.com.br/) (ou similar) para encontrar a localização que deseja consultar a temperatura.
- Utilize a API [WeatherAPI](https://www.weatherapi.com/) (ou similar) para consultar as temperaturas desejadas.
- Para realizar a conversão de Celsius para Fahrenheit, utilize a seguinte fórmula: F = C * 1,8 + 32
- Para realizar a conversão de Celsius para Kelvin, utilize a seguinte fórmula: K = C + 273
  - Sendo: F = Fahrenheit
  - Sendo: C = Celsius
  - Sendo: K = Kelvin

## Entrega
- O código-fonte completo da implementação.
- Testes automatizados demonstrando o funcionamento.
- Utilize docker/docker-compose para que possamos realizar os testes de sua aplicação.
- Deploy realizado no Google Cloud Run (free tier) e endereço ativo para ser acessado.

## Passos para executar o desafio

1. **Clonar o repositório:**

   ```bash
   git clone https://github.com/vinicius-maker/sistema-temperatura-cep.git
   
   ou SSH:
   
   git clone git@github.com:vinicius-maker/sistema-temperatura-cep.git

2. **Configurar o .env:**
       
    ```bash
       cp .env.example .env
   
    Foi necessário colocar outro arquivo .env dentro do diretório cmd para utilização no Cloud RUN (não necessário para testes locais)
   
3. **Configurar o ambiente:**
    ```bash
       docker build -t labs-temperature-cep .
       docker run -d -p 8080:8080 --name labs-temperature-cep -v $(pwd):/app labs-temperature-cep
       docker exec -it labs-temperature-cep bash

4. **Rodar os testes:**
    - dentro do container, em /app:
        ```bash
         go test ./...

5. **Subir a aplicação**
    - dentro do container, em /app/cmd:
        ```bash
        go run main.go
      
6. **Executar aplicação (ambiente local):**
    - url: http://localhost:8080/discover-temperature?cep=<numero-cep>
    - exemplo: http://localhost:8080/discover-temperature?cep=07115260

7. **Executar aplicação (ambiente nuvem):**
    - url: https://labs-temperature-cep-258960345633.us-central1.run.app/discover-temperature?cep=<numero-cep>
    - exemplo: https://labs-temperature-cep-258960345633.us-central1.run.app/discover-temperature?cep=07115260
