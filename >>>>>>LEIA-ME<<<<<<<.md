Hello Dock's dev team !!

Irei escrever em português a partir de agora para mais fácil compreensão :D

A API está escrita na arquitetura clean (Clean Architecture) e está em nuvem: http://177.153.20.221:3001

    Estou compartilhando a collection feita no POSTMAN para maior comodidade, está collection está direcionada para o IP acima.
    Para entender oque foi usado, o README.md e o manifesto.yaml ajudarão.
    Além de estar em nuvem, vocês podem testar local, as instruções estão no README.md.

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