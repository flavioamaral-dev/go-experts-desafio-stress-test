# go-experts-desafio-stress-test
Desafio Pós Go Expert - Stress Test


Sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.


Entrada de Parâmetros via CLI:

--url: URL do serviço a ser testado.
--requests: Número total de requests.
--concurrency: Número de chamadas simultâneas.

Execução do Teste:

- Realizar requests HTTP para a URL especificada.
- Distribuir os requests de acordo com o nível de concorrência definido.
- Garantir que o número total de requests seja cumprido.

- Execução da aplicação:

docker run  —url=https://portal.azure.com —requests=100 —concurrency=10
