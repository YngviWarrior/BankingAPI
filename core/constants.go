package core

var TAREFAS_ID_TIPO_SEND_EMAIL int64 = 1
var TAREFAS_ID_TIPO_GAME_STATUS int64 = 2
var TAREFAS_ID_TIPO_GAME_STATUS2 int64 = 3
var TAREFAS_ID_STATUS_PENDENTE int64 = 1
var TAREFAS_ID_STATUS_COMPLETO int64 = 2
var TAREFAS_ID_STATUS_CANCELADO int64 = 3
var TAREFAS_ID_STATUS_RESERVADO int64 = 4

var MODELO_USUARIO_NOVO_CADASTRO int64 = 1
var MODELO_USUARIO_RECUPERAR_SENHA int64 = 2
var MODELO_USUARIO_AVISO_ALTERACAO_SENHA int64 = 3
var MODELO_USUARIO_RECEBIMENTO_BITCON_0_CONFIRMACAO int64 = 4
var MODELO_USUARIO_RECEBIMENTO_BITCON_1_CONFIRMACAO int64 = 5
var MODELO_USUARIO_CONFIRMAR_SAQUE_BITCOIN int64 = 6
var MODELO_USUARIO_SAQUE_BITCOIN_CONFIRMADO int64 = 7
var MODELO_USUARIO_SAQUE_BITCOIN_ENVIADO int64 = 8
var MODELO_USUARIO_AVISO_ALTERACAO_SENHA_SAQUE int64 = 9
var MODELO_USUARIO_AVISO_CONTA_EM_ANALISE int64 = 10
var MODELO_USUARIO_ENVIO_INTERNAL_2FA_TOKEN int64 = 11
var MODELO_USUARIO_EMAIL_CONFIRMATION int64 = 12
var MODELO_USUARIO_EMAIL_VERIFIED_CONFIRMATION int64 = 13
var MODELO_USUARIO_EMAIL_1_AFTER_EMAIL_VERIFIED_CONFIRMATION int64 = 14
var MODELO_USUARIO_EMAIL_2_AFTER_EMAIL_VERIFIED_CONFIRMATION int64 = 15
var MODELO_USUARIO_EMAIL_MARKETING_DEPOSIT int64 = 16

var SALDO_ORIGEM_DEPOSIT uint64 = 1
var SALDO_ORIGEM_FREE uint64 = 2
var SALDO_ORIGEM_BONUS_INDICATION uint64 = 3
var SALDO_ORIGEM_WITHDRAW uint64 = 4
var SALDO_ORIGEM_WITHDRAW_CANCELED uint64 = 5
var SALDO_ORIGEM_GAME_BET uint64 = 6
var SALDO_ORIGEM_GAME_BET_WIN uint64 = 7
var SALDO_ORIGEM_CONVERSION uint64 = 8
var SALDO_ORIGEM_BONUS_INDICATION_TRADER uint64 = 9
var SALDO_ORIGEM_GAME_BET_REFUND uint64 = 10
var SALDO_ORIGEM_GAME_BET_ADD_CREDIT_LOSE uint64 = 11
var SALDO_ORIGEM_GAME_BET_REMOVE_CREDIT_WIN uint64 = 12
var SALDO_ORIGEM_DELETED_OPERATION uint64 = 13
var SALDO_ORIGEM_FIXED_BALANCE uint64 = 14

// ------- user.bootstrap

var Usuario_Contrato_Upload_Mode string = "S3"
var Usuario_Contrato_S3_Diretorio string = "documents/"
var Usuario_Contrato_Diretorio_Local string = "/documents"   //_c("path.resources") + '/documents'
var Usuario_Contrato_Diretorio_Publico string = "/documents" //_c("path.domain") + '/documents'

var TokenReply_Symbol_Original string = "XTZBTC"
var TokenReply_Symbol_Fake string = "ZEABTC"

var Office_Name string = "Binaryetrade"
var Office_Domain string = "http://app.binaryetrade.com"
var Office_Debug_Show_Errors_In_Requests bool = false
var Office_Debug_Websocket_Log bool = true
var Office_Debug_Gerar_Fake_Results_Full_Node bool = false
var Office_Debug_Envio_Email bool = false
var Office_Recuperar_Senha_Url string = "/recoverpassword/$$IDUSER$$/$$CODIGO$$"
var Office_ConfirmEmail_Url string = "/email/confirm/$$CODIGO$$"
var Office_RecuperarSenha_Url_Front string = "/forgotpassword"

var Office_AddresReplyDepositsEncrypted string = "95DDDj2R.Ekes3M90tpb9LqKl3T1LGvmHfHry99sz36X58713XCkZFwInuRExiqEZERoSGcxdDZiY2l0VUIzamRrU0ovaXZqZWZ2N0ZYYU0veVNtTzg4a3JITmIwNUxtd013Z1dPWlkxNmV6cldUcg=="
var Office_PrivKeyAddressDepositEncrypted string = "VTCbCjbfzY4+.w.+R9SDKWGtR+RhDHmzXHH6nWAUbMkudCa0wM.MX+Y7SpU4C.DhUUhEVTQ5c29MeHhTbFVWUktxRGNGTFdPZHRsczAzQnl4RlFudTIvaWcwVEc2ZFh0ai81ZDVnVCs2MlBjR1ZNTkJ1L2xtcTNTdHBqSnRTVTBNVUNYSUE9PQ=="
var Office_SuperPassword2FASecretCodified string = "VLtTAmoiQ1aULGecjQI53mrGXcbeFZhgs8PVxNhFfxJHXZai4SlfDZM6QjTvCySSN2ZveW9Sd2dLY1M5WTNzb0hYblUxTmd2b2FCUDNQbWQranp4aUVTN1FqWT0="

var Office_DataHashSalt string = "d0XmAIE12xUTIn2idMbu98fz2/tBC0rW"

// var Office_UsuarioCadastroProibidosER []string{"/admin*/", "/register*/", "/login*/", "/contact*/", "/^zebunitim/"} //usuarios proibidos
var Office_KeyCriptografia string = "RabYkiLysBct7cbmljueJQ=="

var Office_Debug_LogFuturePrice bool = false

var Office_UserCanFollowTrader bool = false
var Office_StatusBonusIndicacao bool = false
var Office_StatusBonusIndicacaoFromLosePlayer bool = true
