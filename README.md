<h1>Hello Dock's dev team !!</h1>

A API está dividida em 2 micro-serviços (holder e account),
ambos estão escritos na arquitetura clean (Clean Architecture),
hospedadas em nuvem no IP: http://177.153.20.221.

    O projeto pode ser iniciado localmente, as instruções estão no <b>INSTALL.md</b>.
    Estou compartilhando a API collection criada no POSTMAN para maior comodidade, está collection está direcionada para o IP acima.

Apenas para maior comodidade do avaliador, aqui está um tutorial para fácil manuseio da API.
Sinta-se livre para ignora-lo se for de sua vontade.

    # Gerando o mínimo de informações para uso.

    1. HolderCreate
    2. HolderVerify
    3. AccountCreate
    4. AccountBlock (bloqueio e desbloquei no mesmo endpoint)

    # Teste oque lhe vier na cabeça.

    5. TransactionTypeList (para criar transações, precisamos dos tipos de transações)
    6. AccountTransaction
    7. AccountFind OU HolderFind (nas informações do holder, há uma lista de contas vinculadas a ele)
    8. AccountStatements (Lista as movimentações da conta)

    # Finalizando Conta e "Deletando" o cliente.

    9. AccountDelete (gera um pedido de saque, zerando a conta)
    10. HolderDelete (não acho correto perdermos as informações, então desativei o cliente, NÃO HÁ como ativa-lo novamente pela API)    
    
!!! QUALQUER DÚVIDA, ME LIGUE OU ENVIE UM ZAP (31) 98662-1962 !!!

Atenciosamente, Igor :) 
