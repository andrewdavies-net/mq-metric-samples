package ibmmq

import (
	"strconv"
)

// This file was created from the source tree on 2025-06-10 08:23:17

// This module converts a PCF integer type (from the MQIA* ranges) into
// the prefix that can then be used to convert its value into a string
// via the MQItoString function. For example,
//     prefix:= PCFAttrToPrefix(ibmmq.MQIA_BASE_TYPE)
// returns "MQOT". And then we can call
//     MQItoString(prefix, val) to get "MQOT_Q"
// If there is no appropriate mapping function, and a number is just a number,
// then the returned prefix is the empty string.
// There are no maps for the MQIAMO/MQIAMO64 attributes as those are all raw numbers.
//
// The PCF parameter itself can be turned into the corresponding string by
// calling MQItoString("MQIA",val).
//
// Note that you might still have further work to do if the returned PCF element
// is a LIST (eg MQCFIL) or if the number is a bit-field (like the MQGMO.Options field)
// See the product sample amqsevta.c for ways to deal with those situations.

var pcfAttrMap = map[int32]string{
	MQIA_ACCOUNTING_CONN_OVERRIDE:  "MQMON",
	MQIA_ACCOUNTING_INTERVAL:       "",
	MQIA_ACCOUNTING_MQI:            "MQMON",
	MQIA_ACCOUNTING_Q:              "MQMON",
	MQIA_ACTIVE_CHANNELS:           "",
	MQIA_ACTIVITY_CONN_OVERRIDE:    "MQMON",
	MQIA_ACTIVITY_RECORDING:        "MQRECORDING",
	MQIA_ACTIVITY_TRACE:            "MQMON",
	MQIA_ADOPT_CONTEXT:             "MQADPCTX",
	MQIA_ADOPTNEWMCA_CHECK:         "MQADOPT_CHECK",
	MQIA_ADOPTNEWMCA_INTERVAL:      "",
	MQIA_ADOPTNEWMCA_TYPE:          "MQADOPT_TYPE",
	MQIA_ADVANCED_CAPABILITY:       "MQCAP",
	MQIA_AMQP_CAPABILITY:           "MQCAP",
	MQIA_APPL_TYPE:                 "MQAT",
	MQIA_ARCHIVE:                   "",
	MQIA_AUTHENTICATION_FAIL_DELAY: "",
	MQIA_AUTHENTICATION_METHOD:     "MQAUTHENTICATE",
	MQIA_AUTH_INFO_TYPE:            "MQAIT",
	MQIA_AUTHOREV_SCOPE:            "MQAUSC",
	MQIA_AUTHORITY_EVENT:           "MQEVR",
	MQIA_AUTO_REORGANIZATION:       "MQREORG",
	MQIA_AUTO_REORG_INTERVAL:       "",
	MQIA_BACKOUT_THRESHOLD:         "",
	MQIA_BASE_TYPE:                 "MQOT",
	MQIA_BATCH_INTERFACE_AUTO:      "MQAUTO",
	MQIA_BRIDGE_EVENT:              "MQEVR",
	MQIA_CAP_EXPIRY:                "MQCEX",
	MQIA_CERT_VAL_POLICY:           "MQCERT",
	MQIA_CF_CFCONLOS:               "MQCFCONLOS",
	MQIA_CF_LEVEL:                  "",
	MQIA_CF_OFFLDUSE:               "MQCFOFFLD",
	MQIA_CF_OFFLOAD:                "MQCFOFFLD",
	MQIA_CF_OFFLOAD_THRESHOLD1:     "",
	MQIA_CF_OFFLOAD_THRESHOLD2:     "",
	MQIA_CF_OFFLOAD_THRESHOLD3:     "",
	MQIA_CF_RECAUTO:                "MQRECAUTO",
	MQIA_CF_RECOVER:                "MQCFR",
	MQIA_CF_SMDS_BUFFERS:           "",
	MQIA_CHANNEL_AUTO_DEF:          "MQCHAD",
	MQIA_CHANNEL_AUTO_DEF_EVENT:    "MQEVR",
	MQIA_CHANNEL_EVENT:             "MQEVR",
	MQIA_CHECK_CLIENT_BINDING:      "MQCHK",
	MQIA_CHECK_LOCAL_BINDING:       "MQCHK",
	MQIA_CHINIT_ADAPTERS:           "",
	MQIA_CHINIT_CONTROL:            "MQSVC_CONTROL",
	MQIA_CHINIT_DISPATCHERS:        "",
	MQIA_CHINIT_TRACE_AUTO_START:   "MQTRAXSTR",
	MQIA_CHINIT_TRACE_TABLE_SIZE:   "",
	MQIA_CHLAUTH_RECORDS:           "MQCHLA",
	MQIA_CLUSTER_OBJECT_STATE:      "MQCLST",
	MQIA_CLUSTER_PUB_ROUTE:         "MQCLROUTE",
	MQIA_CLUSTER_Q_TYPE:            "MQCQT",
	MQIA_CLUSTER_WORKLOAD_LENGTH:   "",
	MQIA_CLWL_MRU_CHANNELS:         "",
	MQIA_CLWL_Q_PRIORITY:           "",
	MQIA_CLWL_Q_RANK:               "",
	MQIA_CLWL_USEQ:                 "MQCLWL",
	MQIA_CMD_SERVER_AUTO:           "MQAUTO",
	MQIA_CMD_SERVER_CONTROL:        "MQSVC_CONTROL",
	MQIA_CMD_SERVER_CONVERT_MSG:    "MQCSRV_CONVERT",
	MQIA_CMD_SERVER_DLQ_MSG:        "MQCSRV_DLQ",
	MQIA_CODED_CHAR_SET_ID:         "MQCCSI",
	MQIA_COMMAND_EVENT:             "MQEVR",
	MQIA_COMMAND_LEVEL:             "MQCMDL",
	MQIA_COMM_EVENT:                "MQEVR",
	MQIA_COMM_INFO_TYPE:            "MQCIT",
	MQIA_CONFIGURATION_EVENT:       "MQEVR",
	MQIA_CPI_LEVEL:                 "",
	MQIA_CURRENT_Q_DEPTH:           "",
	MQIA_DEF_BIND:                  "MQBND",
	MQIA_DEF_CLUSTER_XMIT_Q_TYPE:   "MQCLXQ",
	MQIA_DEFINITION_TYPE:           "MQQDT",
	MQIA_DEF_INPUT_OPEN_OPTION:     "MQOO",
	MQIA_DEF_PERSISTENCE:           "MQPER",
	MQIA_DEF_PRIORITY:              "MQPRI",
	MQIA_DEF_PUT_RESPONSE_TYPE:     "MQPRT",
	MQIA_DEF_READ_AHEAD:            "MQREADA",
	MQIA_DISPLAY_TYPE:              "MQDOPT",
	MQIA_DIST_LISTS:                "MQDL",
	MQIA_DNS_WLM:                   "MQDNSWLM",
	MQIA_DURABLE_SUB:               "MQSUB",
	MQIA_ENCRYPTION_ALGORITHM:      "MQMLP_ENCRYPTION",
	MQIA_EXPIRY_INTERVAL:           "MQEXPI",
	MQIA_GROUP_UR:                  "MQGUR",
	MQIA_HARDEN_GET_BACKOUT:        "MQQA_BACKOUT",
	MQIA_HIGH_Q_DEPTH:              "",
	MQIA_IGQ_PUT_AUTHORITY:         "MQIGQPA",
	MQIA_INDEX_TYPE:                "MQIT",
	MQIA_INHIBIT_EVENT:             "MQEVR",
	MQIA_INHIBIT_GET:               "MQQA_GET",
	MQIA_INHIBIT_PUB:               "MQTA_PUB",
	MQIA_INHIBIT_PUT:               "MQQA_PUT",
	MQIA_INHIBIT_SUB:               "MQTA_SUB",
	MQIA_INTRA_GROUP_QUEUING:       "MQIGQ",
	MQIA_IP_ADDRESS_VERSION:        "MQIPADDR",
	MQIA_KEY_REUSE_COUNT:           "MQKEY",
	MQIA_LDAP_AUTHORMD:             "MQLDAP_AUTHORMD",
	MQIA_LDAP_NESTGRP:              "MQLDAP_NESTGRP",
	MQIA_LDAP_SECURE_COMM:          "MQSECCOMM",
	MQIA_LISTENER_PORT_NUMBER:      "",
	MQIA_LISTENER_TIMER:            "",
	MQIA_LOCAL_EVENT:               "MQEVR",
	MQIA_LOGGER_EVENT:              "MQEVR",
	MQIA_LU62_CHANNELS:             "",
	MQIA_MASTER_ADMIN:              "MQMASTER",
	MQIA_MAX_CHANNELS:              "",
	MQIA_MAX_CLIENTS:               "",
	MQIA_MAX_GLOBAL_LOCKS:          "",
	MQIA_MAX_HANDLES:               "",
	MQIA_MAX_LOCAL_LOCKS:           "",
	MQIA_MAX_MSG_LENGTH:            "",
	MQIA_MAX_OPEN_Q:                "",
	MQIA_MAX_PRIORITY:              "",
	MQIA_MAX_PROPERTIES_LENGTH:     "MQPROP",
	MQIA_MAX_Q_DEPTH:               "",
	MQIA_MAX_Q_FILE_SIZE:           "MQQFS",
	MQIA_MAX_Q_TRIGGERS:            "",
	MQIA_MAX_RECOVERY_TASKS:        "",
	MQIA_MAX_RESPONSES:             "",
	MQIA_MAX_UNCOMMITTED_MSGS:      "",
	MQIA_MCAST_BRIDGE:              "MQMCB",
	MQIA_MEDIA_IMAGE_INTERVAL:      "MQMEDIMGINTVL",
	MQIA_MEDIA_IMAGE_LOG_LENGTH:    "MQMEDIMGLOGLN",
	MQIA_MEDIA_IMAGE_RECOVER_OBJ:   "MQIMGRCOV",
	MQIA_MEDIA_IMAGE_RECOVER_Q:     "MQIMGRCOV",
	MQIA_MEDIA_IMAGE_SCHEDULING:    "MQMEDIMGSCHED",
	MQIA_MONITORING_AUTO_CLUSSDR:   "MQMON",
	MQIA_MONITORING_CHANNEL:        "MQMON",
	MQIA_MONITORING_Q:              "MQMON",
	MQIA_MONITOR_INTERVAL:          "",
	MQIA_MSG_DELIVERY_SEQUENCE:     "MQMDS",
	MQIA_MSG_DEQ_COUNT:             "",
	MQIA_MSG_ENQ_COUNT:             "",
	MQIA_MSG_MARK_BROWSE_INTERVAL:  "MQMMBI",
	MQIA_MULTICAST:                 "MQMC",
	MQIA_NAME_COUNT:                "",
	MQIA_NAMELIST_TYPE:             "MQNT",
	MQIA_NPM_CLASS:                 "MQNPM",
	MQIA_NPM_DELIVERY:              "MQDLV",
	MQIA_OPEN_INPUT_COUNT:          "",
	MQIA_OPEN_OUTPUT_COUNT:         "",
	MQIA_OTEL_PROPAGATION_CONTROL:  "MQOTEL_PCTL",
	MQIA_OTEL_TRACE:                "MQOTEL_TRACE",
	MQIA_OUTBOUND_PORT_MAX:         "",
	MQIA_OUTBOUND_PORT_MIN:         "",
	MQIA_PAGESET_ID:                "",
	MQIA_PERFORMANCE_EVENT:         "MQEVR",
	MQIA_PLATFORM:                  "MQPL",
	MQIA_PM_DELIVERY:               "MQDLV",
	MQIA_POLICY_VERSION:            "",
	MQIA_PROPERTY_CONTROL:          "MQPROP",
	MQIA_PROT_POLICY_CAPABILITY:    "MQCAP",
	MQIA_PROXY_SUB:                 "MQTA_PROXY",
	MQIA_PUB_COUNT:                 "",
	MQIA_PUB_SCOPE:                 "MQSCOPE",
	MQIA_PUBSUB_CLUSTER:            "MQPSCLUS",
	MQIA_PUBSUB_MAXMSG_RETRY_COUNT: "",
	MQIA_PUBSUB_MODE:               "MQPSM",
	MQIA_PUBSUB_NP_MSG:             "MQUNDELIVERED",
	MQIA_PUBSUB_NP_RESP:            "MQUNDELIVERED",
	MQIA_PUBSUB_SYNC_PT:            "MQSYNCPOINT",
	MQIA_Q_DEPTH_HIGH_EVENT:        "MQEVR",
	MQIA_Q_DEPTH_HIGH_LIMIT:        "",
	MQIA_Q_DEPTH_LOW_EVENT:         "MQEVR",
	MQIA_Q_DEPTH_LOW_LIMIT:         "",
	MQIA_Q_DEPTH_MAX_EVENT:         "MQEVR",
	MQIA_QMGR_CFCONLOS:             "MQCFCONLOS",
	MQIA_QMOPT_CONS_COMMS_MSGS:     "MQQMOPT",
	MQIA_QMOPT_CONS_CRITICAL_MSGS:  "MQQMOPT",
	MQIA_QMOPT_CONS_ERROR_MSGS:     "MQQMOPT",
	MQIA_QMOPT_CONS_INFO_MSGS:      "MQQMOPT",
	MQIA_QMOPT_CONS_REORG_MSGS:     "MQQMOPT",
	MQIA_QMOPT_CONS_SYSTEM_MSGS:    "MQQMOPT",
	MQIA_QMOPT_CONS_WARNING_MSGS:   "MQQMOPT",
	MQIA_QMOPT_CSMT_ON_ERROR:       "MQQMOPT",
	MQIA_QMOPT_INTERNAL_DUMP:       "MQQMOPT",
	MQIA_QMOPT_LOG_COMMS_MSGS:      "MQQMOPT",
	MQIA_QMOPT_LOG_CRITICAL_MSGS:   "MQQMOPT",
	MQIA_QMOPT_LOG_ERROR_MSGS:      "MQQMOPT",
	MQIA_QMOPT_LOG_INFO_MSGS:       "MQQMOPT",
	MQIA_QMOPT_LOG_REORG_MSGS:      "MQQMOPT",
	MQIA_QMOPT_LOG_SYSTEM_MSGS:     "MQQMOPT",
	MQIA_QMOPT_LOG_WARNING_MSGS:    "MQQMOPT",
	MQIA_QMOPT_TRACE_COMMS:         "MQQMOPT",
	MQIA_QMOPT_TRACE_CONVERSION:    "MQQMOPT",
	MQIA_QMOPT_TRACE_MQI_CALLS:     "MQQMOPT",
	MQIA_QMOPT_TRACE_REORG:         "MQQMOPT",
	MQIA_QMOPT_TRACE_SYSTEM:        "MQQMOPT",
	MQIA_Q_SERVICE_INTERVAL:        "",
	MQIA_Q_SERVICE_INTERVAL_EVENT:  "MQQSIE",
	MQIA_QSG_DISP:                  "MQQSGD",
	MQIA_Q_TYPE:                    "MQQT",
	MQIA_Q_USERS:                   "",
	MQIA_READ_AHEAD:                "MQREADA",
	MQIA_RECEIVE_TIMEOUT:           "",
	MQIA_RECEIVE_TIMEOUT_MIN:       "",
	MQIA_RECEIVE_TIMEOUT_TYPE:      "MQRCVTIME",
	MQIA_REMOTE_EVENT:              "MQEVR",
	MQIA_RESPONSE_RESTART_POINT:    "",
	MQIA_RETENTION_INTERVAL:        "",
	MQIA_REVERSE_DNS_LOOKUP:        "MQRDNS",
	MQIA_SCOPE:                     "MQSCO",
	MQIA_SECURITY_CASE:             "MQSCYC",
	MQIA_SERVICE_CONTROL:           "MQSVC_CONTROL",
	MQIA_SERVICE_TYPE:              "MQSVC_TYPE",
	MQIA_SHAREABILITY:              "MQQA_SHAREABLE",
	MQIA_SHARED_Q_Q_MGR_NAME:       "MQSQQM",
	MQIA_SIGNATURE_ALGORITHM:       "MQMLP_SIGN",
	MQIA_SSL_EVENT:                 "MQEVR",
	MQIA_SSL_FIPS_REQUIRED:         "MQSSL",
	MQIA_SSL_RESET_COUNT:           "",
	MQIA_SSL_TASKS:                 "",
	MQIA_START_STOP_EVENT:          "MQEVR",
	MQIA_STATISTICS_AUTO_CLUSSDR:   "MQMON",
	MQIA_STATISTICS_CHANNEL:        "MQMON",
	MQIA_STATISTICS_INTERVAL:       "",
	MQIA_STATISTICS_MQI:            "MQMON",
	MQIA_STATISTICS_Q:              "MQMON",
	MQIA_STREAM_QUEUE_QOS:          "MQST",
	MQIA_SUB_CONFIGURATION_EVENT:   "MQEVR",
	MQIA_SUB_COUNT:                 "MQPSCT",
	MQIA_SUB_SCOPE:                 "MQSCOPE",
	MQIA_SUITE_B_STRENGTH:          "MQSUITE",
	MQIA_SYNCPOINT:                 "MQSP",
	MQIA_TCP_CHANNELS:              "",
	MQIA_TCP_KEEP_ALIVE:            "MQTCPKEEP",
	MQIA_TCP_STACK_TYPE:            "MQTCPSTACK",
	MQIA_TIME_SINCE_RESET:          "",
	MQIA_TOLERATE_UNPROTECTED:      "MQMLP_TOLERATE",
	MQIA_TOPIC_DEF_PERSISTENCE:     "MQPER",
	MQIA_TOPIC_NODE_COUNT:          "MQPSCT",
	MQIA_TOPIC_TYPE:                "MQTOPT",
	MQIA_TRACE_ROUTE_RECORDING:     "MQRECORDING",
	MQIA_TREE_LIFE_TIME:            "",
	MQIA_TRIGGER_CONTROL:           "MQTC",
	MQIA_TRIGGER_DEPTH:             "",
	MQIA_TRIGGER_INTERVAL:          "",
	MQIA_TRIGGER_MSG_PRIORITY:      "",
	MQIA_TRIGGER_RESTART:           "MQTRIGGER",
	MQIA_TRIGGER_TYPE:              "MQTT",
	MQIA_UR_DISP:                   "MQQSGD",
	MQIA_USAGE:                     "MQUS",
	MQIA_USE_DEAD_LETTER_Q:         "MQUSEDLQ",
	MQIA_USER_LIST:                 "",
	MQIA_WILDCARD_OPERATION:        "MQTA",
	MQIA_XR_CAPABILITY:             "MQCAP",
	MQIACF_ACTION:                  "MQACT",
	MQIACF_ALL:                     "",
	MQIACF_AMQP_DIAGNOSTICS_TYPE:   "",
	MQIACF_ANONYMOUS_COUNT:         "",
	MQIACF_API_CALLER_TYPE:         "MQXACT",
	MQIACF_API_ENVIRONMENT:         "MQXE",
	MQIACF_APPL_COUNT:              "",
	MQIACF_APPL_FUNCTION_TYPE:      "MQFUN",
	MQIACF_APPL_IMMOVABLE_COUNT:    "",
	MQIACF_APPL_IMMOVABLE_REASON:   "MQIMMREASON",
	MQIACF_APPL_INFO_APPL:          "MQAPPL",
	MQIACF_APPL_INFO_LOCAL:         "",
	MQIACF_APPL_INFO_QMGR:          "",
	MQIACF_APPL_INFO_TYPE:          "",
	MQIACF_APPL_MOVABLE:            "MQACTIVE",
	MQIACF_ARCHIVE_LOG_SIZE:        "",
	MQIACF_ASYNC_STATE:             "MQAS",
	MQIACF_AUTH_ADD_AUTHS:          "MQAUTH",
	MQIACF_AUTH_OPTIONS:            "MQAUTHOPT",
	MQIACF_AUTHORIZATION_LIST:      "MQAUTH",
	MQIACF_AUTH_REC_TYPE:           "MQOT",
	MQIACF_AUTH_REMOVE_AUTHS:       "MQAUTH",
	MQIACF_AUTO_CLUSTER_TYPE:       "MQAUTOCLUS",
	MQIACF_AUX_ERROR_DATA_INT_1:    "",
	MQIACF_AUX_ERROR_DATA_INT_2:    "",
	MQIACF_BACKOUT_COUNT:           "",
	MQIACF_BALANCED:                "MQBALANCED",
	MQIACF_BALANCING_OPTIONS:       "MQBNO_OPTIONS",
	MQIACF_BALANCING_TIMEOUT:       "MQBNO_TIMEOUT",
	MQIACF_BALANCING_TYPE:          "MQBNO_BALTYPE",
	MQIACF_BALSTATE:                "MQBALSTATE",
	MQIACF_BRIDGE_TYPE:             "MQBT",
	MQIACF_BROKER_COUNT:            "",
	MQIACF_BROKER_OPTIONS:          "",
	MQIACF_BUFFER_LENGTH:           "",
	MQIACF_BUFFER_POOL_ID:          "",
	MQIACF_BUFFER_POOL_LOCATION:    "MQBPLOCATION",
	MQIACF_CALL_TYPE:               "MQCBCT",
	MQIACF_CF_SMDS_BLOCK_SIZE:      "MQDSB",
	MQIACF_CF_SMDS_EXPAND:          "MQDSE",
	MQIACF_CF_STATUS_BACKUP:        "MQCFSTATUS",
	MQIACF_CF_STATUS_CONNECT:       "MQCFSTATUS",
	MQIACF_CF_STATUS_SMDS:          "MQCFSTATUS",
	MQIACF_CF_STATUS_SUMMARY:       "MQCFSTATUS",
	MQIACF_CF_STATUS_TYPE:          "MQCFSTATUS",
	MQIACF_CF_STRUC_ACCESS:         "MQCFACCESS",
	MQIACF_CF_STRUC_BACKUP_SIZE:    "",
	MQIACF_CF_STRUC_ENTRIES_MAX:    "",
	MQIACF_CF_STRUC_ENTRIES_USED:   "",
	MQIACF_CF_STRUC_SIZE_MAX:       "",
	MQIACF_CF_STRUC_SIZE_USED:      "",
	MQIACF_CF_STRUC_STATUS:         "MQCFSTATUS",
	MQIACF_CF_STRUC_TYPE:           "MQCFTYPE",
	MQIACF_CHECKPOINT_COUNT:        "",
	MQIACF_CHECKPOINT_OPERATIONS:   "",
	MQIACF_CHECKPOINT_SIZE:         "",
	MQIACF_CHINIT_STATUS:           "MQSVC_STATUS",
	MQIACF_CHLAUTH_TYPE:            "MQCAUT",
	MQIACF_CLEAR_SCOPE:             "MQCLRS",
	MQIACF_CLEAR_TYPE:              "MQCLRT",
	MQIACF_CLOSE_OPTIONS:           "MQCO",
	MQIACF_CLUSTER_INFO:            "",
	MQIACF_CMDSCOPE_Q_MGR_COUNT:    "",
	MQIACF_CMD_SERVER_STATUS:       "MQSVC_CONTROL",
	MQIACF_COMMAND:                 "MQCMD",
	MQIACF_COMMAND_INFO:            "MQCMDI",
	MQIACF_COMP_CODE:               "MQCC",
	MQIACF_CONFIGURATION_EVENTS:    "",
	MQIACF_CONFIGURATION_OBJECTS:   "",
	MQIACF_CONNECTION_COUNT:        "",
	MQIACF_CONNECTION_SWAP:         "",
	MQIACF_CONNECT_OPTIONS:         "MQCNO",
	MQIACF_CONNECT_TIME:            "",
	MQIACF_CONN_INFO_ALL:           "",
	MQIACF_CONN_INFO_CONN:          "",
	MQIACF_CONN_INFO_HANDLE:        "",
	MQIACF_CONN_INFO_TYPE:          "",
	MQIACF_CONV_REASON_CODE:        "MQRC",
	MQIACF_CTL_OPERATION:           "MQOP",
	MQIACF_CUR_MAX_FILE_SIZE:       "",
	MQIACF_CUR_Q_FILE_SIZE:         "",
	MQIACF_DATA_FS_IN_USE:          "MQFS",
	MQIACF_DATA_FS_SIZE:            "MQFS",
	MQIACF_DB2_CONN_STATUS:         "MQQSGS",
	MQIACF_DELETE_OPTIONS:          "MQDMHO",
	MQIACF_DESTINATION_CLASS:       "MQDC",
	MQIACF_DISCONNECT_TIME:         "",
	MQIACF_DISCONTINUITY_COUNT:     "",
	MQIACF_DS_ENCRYPTED:            "MQSYSP",
	MQIACF_DURABLE_SUBSCRIPTION:    "MQSUB",
	MQIACF_ENCODING:                "MQENC",
	MQIACF_ENTITY_TYPE:             "MQZAET",
	MQIACF_ERROR_ID:                "MQRC",
	MQIACF_ERROR_OFFSET:            "",
	MQIACF_ESCAPE_TYPE:             "MQET",
	MQIACF_EVENT_APPL_TYPE:         "MQAT",
	MQIACF_EVENT_DUPLICATE_COUNT:   "",
	MQIACF_EVENT_ORIGIN:            "MQEVO",
	MQIACF_EXCLUDE_INTERVAL:        "",
	MQIACF_EXPIRY:                  "",
	MQIACF_EXPIRY_Q_COUNT:          "",
	MQIACF_EXPIRY_TIME:             "MQEI",
	MQIACF_EXPORT_TYPE:             "MQEXT",
	MQIACF_FEEDBACK:                "MQFB",
	MQIACF_FORCE:                   "MQFC",
	MQIACF_GET_OPTIONS:             "MQGMO",
	MQIACF_GROUPUR_CHECK_ID:        "",
	MQIACF_HANDLE_STATE:            "MQHSTATE",
	MQIACF_HOBJ:                    "",
	MQIACF_HSUB:                    "",
	MQIACF_IGNORE_STATE:            "MQIS",
	MQIACF_INQUIRY:                 "",
	MQIACF_INTATTR_COUNT:           "",
	MQIACF_INTEGER_DATA:            "",
	MQIACF_INTERFACE_VERSION:       "",
	MQIACF_INVALID_DEST_COUNT:      "",
	MQIACF_ITEM_COUNT:              "",
	MQIACF_KNOWN_DEST_COUNT:        "",
	MQIACF_LDAP_CONNECTION_STATUS:  "MQLDAPC",
	MQIACF_LOG_COMPRESSION:         "MQCOMPRESS",
	MQIACF_LOG_EXTENT_SIZE:         "",
	MQIACF_LOG_FS_IN_USE:           "MQFS",
	MQIACF_LOG_FS_SIZE:             "MQFS",
	MQIACF_LOG_IN_USE:              "",
	MQIACF_LOG_PRIMARIES:           "",
	MQIACF_LOG_REDUCTION:           "MQLR",
	MQIACF_LOG_SECONDARIES:         "",
	MQIACF_LOG_TYPE:                "MQLOGTYPE",
	MQIACF_LOG_UTILIZATION:         "",
	MQIACF_MAX_ACTIVITIES:          "MQROUTE",
	MQIACF_MCAST_REL_INDICATOR:     "",
	MQIACF_MEDIA_LOG_SIZE:          "",
	MQIACF_MESSAGE_COUNT:           "",
	MQIACF_MODE:                    "MQMODE",
	MQIACF_MONITORING:              "MQMON",
	MQIACF_MOVABLE_APPL_COUNT:      "",
	MQIACF_MOVE_COUNT:              "",
	MQIACF_MOVE_TYPE:               "",
	MQIACF_MOVE_TYPE_ADD:           "",
	MQIACF_MOVE_TYPE_MOVE:          "",
	MQIACF_MQCB_OPERATION:          "MQOP",
	MQIACF_MQCB_OPTIONS:            "MQCBDO",
	MQIACF_MQCB_TYPE:               "MQCBT",
	MQIACF_MQXR_DIAGNOSTICS_TYPE:   "",
	MQIACF_MSG_FLAGS:               "MQMF",
	MQIACF_MSG_LENGTH:              "",
	MQIACF_MSG_TYPE:                "MQMT",
	MQIACF_MULC_CAPTURE:            "MQMULC",
	MQIACF_NHA_GROUP_BACKLOG:       "MQNHABACKLOG",
	MQIACF_NHA_GROUP_CONNECTED:     "MQNHACONNGRP",
	MQIACF_NHA_GROUP_IN_SYNC:       "MQNHAINSYNC",
	MQIACF_NHA_GROUP_ROLE:          "MQNHAGRPROLE",
	MQIACF_NHA_GROUP_STATUS:        "MQNHASTATUS",
	MQIACF_NHA_INSTANCE_ACTV_CONNS: "MQNHACONNACTV",
	MQIACF_NHA_INSTANCE_BACKLOG:    "MQNHABACKLOG",
	MQIACF_NHA_INSTANCE_IN_SYNC:    "MQNHAINSYNC",
	MQIACF_NHA_INSTANCE_ROLE:       "MQNHAROLE",
	MQIACF_NHA_INSTANCE_STATUS:     "MQNHASTATUS",
	MQIACF_NHA_IN_SYNC_INSTANCES:   "",
	MQIACF_NHA_TOTAL_INSTANCES:     "",
	MQIACF_NHA_TYPE:                "MQNHATYPE",
	MQIACF_NUM_PUBS:                "",
	MQIACF_OBJECT_TYPE:             "MQOT",
	MQIACF_OBSOLETE_MSGS:           "MQOM",
	MQIACF_OFFSET:                  "",
	MQIACF_OLDEST_MSG_AGE:          "",
	MQIACF_OPEN_BROWSE:             "MQQSO",
	MQIACF_OPEN_INPUT_TYPE:         "MQQSO",
	MQIACF_OPEN_INQUIRE:            "MQQSO",
	MQIACF_OPEN_OPTIONS:            "MQOO",
	MQIACF_OPEN_OUTPUT:             "MQQSO",
	MQIACF_OPEN_SET:                "MQQSO",
	MQIACF_OPEN_TYPE:               "MQQSOT",
	MQIACF_OPERATION_ID:            "MQXF",
	MQIACF_OPERATION_MODE:          "MQOPMODE",
	MQIACF_OPERATION_TYPE:          "MQOPER",
	MQIACF_OPTIONS:                 "",
	MQIACF_ORIGINAL_LENGTH:         "MQOL",
	MQIACF_PAGECLAS:                "MQPAGECLAS",
	MQIACF_PAGESET_STATUS:          "MQUSAGE_PS",
	MQIACF_PARAMETER_ID:            "",
	MQIACF_PERMIT_STANDBY:          "MQSTDBY",
	MQIACF_PERSISTENCE:             "MQPER",
	MQIACF_POINTER_SIZE:            "",
	MQIACF_PRIORITY:                "MQPRI",
	MQIACF_PROCESS_ID:              "",
	MQIACF_PS_STATUS_TYPE:          "MQPSST",
	MQIACF_PUBLICATION_OPTIONS:     "MQPUBO",
	MQIACF_PUBLISH_COUNT:           "",
	MQIACF_PUB_PRIORITY:            "MQPRI",
	MQIACF_PUBSUB_PROPERTIES:       "MQPSPROP",
	MQIACF_PUBSUB_STATUS:           "MQPS",
	MQIACF_PURGE:                   "MQPO",
	MQIACF_PUT_OPTIONS:             "MQPMO",
	MQIACF_Q_HANDLE:                "",
	MQIACF_Q_MGR_CLUSTER:           "",
	MQIACF_Q_MGR_DEFINITION_TYPE:   "MQQMDT",
	MQIACF_Q_MGR_DQM:               "",
	MQIACF_Q_MGR_EVENT:             "",
	MQIACF_Q_MGR_FACILITY:          "MQQMFAC",
	MQIACF_Q_MGR_FS_ENCRYPTED:      "MQFSENC",
	MQIACF_Q_MGR_FS_IN_USE:         "",
	MQIACF_Q_MGR_FS_SIZE:           "",
	MQIACF_Q_MGR_NUMBER:            "",
	MQIACF_Q_MGR_PUBSUB:            "MQPSM",
	MQIACF_Q_MGR_STATUS:            "MQQMSTA",
	MQIACF_Q_MGR_STATUS_INFO_NHA:   "",
	MQIACF_Q_MGR_STATUS_INFO_Q_MGR: "",
	MQIACF_Q_MGR_STATUS_INFO_TYPE:  "",
	MQIACF_Q_MGR_STATUS_LOG:        "",
	MQIACF_Q_MGR_SYSTEM:            "",
	MQIACF_Q_MGR_TYPE:              "MQQMT",
	MQIACF_Q_MGR_VERSION:           "",
	MQIACF_QSG_DISPS:               "MQQSGD",
	MQIACF_Q_STATUS:                "",
	MQIACF_Q_STATUS_TYPE:           "",
	MQIACF_Q_TIME_INDICATOR:        "MQMON",
	MQIACF_Q_TYPES:                 "MQQT",
	MQIACF_REASON_CODE:             "MQRC",
	MQIACF_REASON_QUALIFIER:        "MQRQ",
	MQIACF_RECORDED_ACTIVITIES:     "",
	MQIACF_RECS_PRESENT:            "",
	MQIACF_REFRESH_INTERVAL:        "",
	MQIACF_REFRESH_REPOSITORY:      "MQCFO_REFRESH",
	MQIACF_REFRESH_TYPE:            "MQRT",
	MQIACF_REGISTRATION_OPTIONS:    "MQREGO",
	MQIACF_REG_REG_OPTIONS:         "",
	MQIACF_REMOTE_QMGR_ACTIVE:      "MQACTIVE",
	MQIACF_REMOVE_AUTHREC:          "MQRAR",
	MQIACF_REMOVE_QUEUES:           "MQCFO_REMOVE",
	MQIACF_REPLACE:                 "MQRP",
	MQIACF_REPORT:                  "MQRO",
	MQIACF_REQUEST_ONLY:            "MQRU",
	MQIACF_RESOLVED_TYPE:           "MQOT",
	MQIACF_RESTART_LOG_SIZE:        "",
	MQIACF_RETAINED_PUBLICATION:    "MQQSO",
	MQIACF_REUSABLE_LOG_SIZE:       "",
	MQIACF_ROUTE_ACCUMULATION:      "MQROUTE",
	MQIACF_ROUTE_DELIVERY:          "MQROUTE",
	MQIACF_ROUTE_DETAIL:            "MQROUTE",
	MQIACF_ROUTE_FORWARDING:        "MQROUTE",
	MQIACF_SECURITY_INTERVAL:       "",
	MQIACF_SECURITY_ITEM:           "MQSECITEM",
	MQIACF_SECURITY_SETTING:        "MQSECSW",
	MQIACF_SECURITY_SWITCH:         "MQSECSW",
	MQIACF_SECURITY_TIMEOUT:        "",
	MQIACF_SECURITY_TYPE:           "MQSECTYPE",
	MQIACF_SELECTOR:                "",
	MQIACF_SELECTOR_COUNT:          "",
	MQIACF_SELECTORS:               "",
	MQIACF_SELECTOR_TYPE:           "MQSELTYPE",
	MQIACF_SEQUENCE_NUMBER:         "",
	MQIACF_SERVICE_STATUS:          "MQSVC_STATUS",
	MQIACF_SMDS_AVAIL:              "MQS_AVAIL",
	MQIACF_SMDS_EXPANDST:           "MQS_EXPANDST",
	MQIACF_SMDS_OPENMODE:           "MQS_OPENMODE",
	MQIACF_SMDS_STATUS:             "MQUSAGE_SMDS",
	MQIACF_STATUS_TYPE:             "MQSTAT",
	MQIACF_STRUC_LENGTH:            "",
	MQIACF_SUB_LEVEL:               "",
	MQIACF_SUB_OPTIONS:             "MQSO",
	MQIACF_SUBRQ_ACTION:            "MQSR",
	MQIACF_SUBRQ_OPTIONS:           "MQSRO",
	MQIACF_SUBSCRIPTION_SCOPE:      "MQTSCOPE",
	MQIACF_SUB_SUMMARY:             "",
	MQIACF_SUB_TYPE:                "MQSUBTYPE",
	MQIACF_SUSPEND:                 "MQSUS",
	MQIACF_SYSP_ALLOC_PRIMARY:      "",
	MQIACF_SYSP_ALLOC_SECONDARY:    "",
	MQIACF_SYSP_ALLOC_UNIT:         "MQSYSP",
	MQIACF_SYSP_ARCHIVE:            "MQSYSP",
	MQIACF_SYSP_ARCHIVE_RETAIN:     "MQSYSP",
	MQIACF_SYSP_ARCHIVE_WTOR:       "MQSYSP",
	MQIACF_SYSP_BLOCK_SIZE:         "",
	MQIACF_SYSP_CATALOG:            "MQSYSP",
	MQIACF_SYSP_CHKPOINT_COUNT:     "",
	MQIACF_SYSP_CLUSTER_CACHE:      "MQCLCT",
	MQIACF_SYSP_COMPACT:            "MQSYSP",
	MQIACF_SYSP_DB2_BLOB_TASKS:     "",
	MQIACF_SYSP_DB2_TASKS:          "",
	MQIACF_SYSP_DEALLOC_INTERVAL:   "",
	MQIACF_SYSP_DUAL_ACTIVE:        "MQSYSP",
	MQIACF_SYSP_DUAL_ARCHIVE:       "MQSYSP",
	MQIACF_SYSP_DUAL_BSDS:          "MQSYSP",
	MQIACF_SYSP_EXIT_INTERVAL:      "",
	MQIACF_SYSP_EXIT_TASKS:         "",
	MQIACF_SYSP_FULL_LOGS:          "",
	MQIACF_SYSP_IN_BUFFER_SIZE:     "",
	MQIACF_SYSP_LOG_COPY:           "",
	MQIACF_SYSP_LOG_SUSPEND:        "MQSYSP",
	MQIACF_SYSP_LOG_USED:           "",
	MQIACF_SYSP_MAX_ACE_POOL:       "",
	MQIACF_SYSP_MAX_ARCHIVE:        "",
	MQIACF_SYSP_MAX_CONC_OFFLOADS:  "",
	MQIACF_SYSP_MAX_CONNS:          "",
	MQIACF_SYSP_MAX_CONNS_BACK:     "",
	MQIACF_SYSP_MAX_CONNS_FORE:     "",
	MQIACF_SYSP_MAX_READ_TAPES:     "",
	MQIACF_SYSP_OFFLOAD_STATUS:     "MQSYSP",
	MQIACF_SYSP_OTMA_INTERVAL:      "",
	MQIACF_SYSP_OUT_BUFFER_COUNT:   "",
	MQIACF_SYSP_OUT_BUFFER_SIZE:    "",
	MQIACF_SYSP_PROTECT:            "MQSYSP",
	MQIACF_SYSP_Q_INDEX_DEFER:      "MQSYSP",
	MQIACF_SYSP_QUIESCE_INTERVAL:   "",
	MQIACF_SYSP_RESLEVEL_AUDIT:     "MQSYSP",
	MQIACF_SYSP_ROUTING_CODE:       "",
	MQIACF_SYSP_SMF_ACCOUNTING:     "MQSYSP",
	MQIACF_SYSP_SMF_ACCT_TIME_MINS: "",
	MQIACF_SYSP_SMF_ACCT_TIME_SECS: "",
	MQIACF_SYSP_SMF_STATS:          "MQSYSP",
	MQIACF_SYSP_SMF_STAT_TIME_MINS: "",
	MQIACF_SYSP_SMF_STAT_TIME_SECS: "",
	MQIACF_SYSP_TIMESTAMP:          "MQSYSP",
	MQIACF_SYSP_TOTAL_LOGS:         "",
	MQIACF_SYSP_TRACE_CLASS:        "",
	MQIACF_SYSP_TRACE_SIZE:         "",
	MQIACF_SYSP_TYPE:               "MQSYSP",
	MQIACF_SYSP_UNIT_ADDRESS:       "",
	MQIACF_SYSP_UNIT_STATUS:        "MQSYSP",
	MQIACF_SYSP_WLM_INTERVAL:       "",
	MQIACF_SYSP_WLM_INT_UNITS:      "MQTIME",
	MQIACF_SYSP_ZHYPERLINK:         "MQSYSP",
	MQIACF_SYSP_ZHYPERWRITE:        "MQSYSP",
	MQIACF_SYSTEM_OBJECTS:          "MQSYSOBJ",
	MQIACF_THREAD_ID:               "",
	MQIACF_TOPIC_PUB:               "",
	MQIACF_TOPIC_STATUS:            "",
	MQIACF_TOPIC_STATUS_TYPE:       "",
	MQIACF_TOPIC_SUB:               "",
	MQIACF_TRACE_DATA_LENGTH:       "",
	MQIACF_TRACE_DETAIL:            "MQACTV",
	MQIACF_UNCOMMITTED_MSGS:        "",
	MQIACF_UNKNOWN_DEST_COUNT:      "",
	MQIACF_UNRECORDED_ACTIVITIES:   "",
	MQIACF_UOW_STATE:               "MQUOWST",
	MQIACF_UOW_TYPE:                "MQUOWT",
	MQIACF_USAGE_BLOCK_SIZE:        "",
	MQIACF_USAGE_BUFFER_POOL:       "",
	MQIACF_USAGE_DATA_BLOCKS:       "",
	MQIACF_USAGE_DATA_SET:          "",
	MQIACF_USAGE_DATA_SET_TYPE:     "MQUSAGE_DS",
	MQIACF_USAGE_EMPTY_BUFFERS:     "",
	MQIACF_USAGE_EXPAND_COUNT:      "MQUSAGE_EXPAND",
	MQIACF_USAGE_EXPAND_TYPE:       "MQUSAGE_EXPAND",
	MQIACF_USAGE_FREE_BUFF:         "",
	MQIACF_USAGE_FREE_BUFF_PERC:    "",
	MQIACF_USAGE_INUSE_BUFFERS:     "",
	MQIACF_USAGE_LOWEST_FREE:       "",
	MQIACF_USAGE_NONPERSIST_PAGES:  "",
	MQIACF_USAGE_OFFLOAD_MSGS:      "",
	MQIACF_USAGE_PAGESET:           "",
	MQIACF_USAGE_PERSIST_PAGES:     "",
	MQIACF_USAGE_READS_SAVED:       "",
	MQIACF_USAGE_RESTART_EXTENTS:   "",
	MQIACF_USAGE_SAVED_BUFFERS:     "",
	MQIACF_USAGE_SMDS:              "",
	MQIACF_USAGE_TOTAL_BLOCKS:      "",
	MQIACF_USAGE_TOTAL_BUFFERS:     "",
	MQIACF_USAGE_TOTAL_PAGES:       "",
	MQIACF_USAGE_TYPE:              "",
	MQIACF_USAGE_UNUSED_PAGES:      "",
	MQIACF_USAGE_USED_BLOCKS:       "",
	MQIACF_USAGE_USED_RATE:         "",
	MQIACF_USAGE_WAIT_RATE:         "",
	MQIACF_USER_ID_SUPPORT:         "MQUIDSUPP",
	MQIACF_VARIABLE_USER_ID:        "MQVU",
	MQIACF_VERSION:                 "",
	MQIACF_WAIT_INTERVAL:           "MQWI",
	MQIACF_WILDCARD_SCHEMA:         "MQWS",
	MQIACF_XA_COUNT:                "",
	MQIACF_XA_FLAGS:                "",
	MQIACF_XA_HANDLE:               "",
	MQIACF_XA_RETCODE:              "",
	MQIACF_XA_RETVAL:               "",
	MQIACF_XA_RMID:                 "",
	MQIACH_ACTIVE_CHL:              "",
	MQIACH_ACTIVE_CHL_MAX:          "",
	MQIACH_ACTIVE_CHL_PAUSED:       "",
	MQIACH_ACTIVE_CHL_RETRY:        "",
	MQIACH_ACTIVE_CHL_STARTED:      "",
	MQIACH_ACTIVE_CHL_STOPPED:      "",
	MQIACH_ADAPS_MAX:               "",
	MQIACH_ADAPS_STARTED:           "",
	MQIACH_ADAPTER:                 "",
	MQIACH_ALLOC_FAST_TIMER:        "",
	MQIACH_ALLOC_RETRY:             "",
	MQIACH_ALLOC_SLOW_TIMER:        "",
	MQIACH_AMQP_KEEP_ALIVE:         "MQKAI",
	MQIACH_AUTH_INFO_TYPES:         "MQAIT",
	MQIACH_AVAILABLE_CIPHERSPECS:   "",
	MQIACH_BACKLOG:                 "",
	MQIACH_BATCH_DATA_LIMIT:        "",
	MQIACH_BATCHES:                 "",
	MQIACH_BATCH_HB:                "",
	MQIACH_BATCH_INTERVAL:          "",
	MQIACH_BATCH_SIZE:              "",
	MQIACH_BATCH_SIZE_INDICATOR:    "",
	MQIACH_BUFFERS_RCVD:            "",
	MQIACH_BUFFERS_SENT:            "",
	MQIACH_BYTES_RCVD:              "",
	MQIACH_BYTES_SENT:              "",
	MQIACH_CHANNEL_DISP:            "MQCHLD",
	MQIACH_CHANNEL_ERROR_DATA:      "",
	MQIACH_CHANNEL_INSTANCE_TYPE:   "MQOT",
	MQIACH_CHANNEL_STATUS:          "MQCHS",
	MQIACH_CHANNEL_SUBSTATE:        "MQCHSSTATE",
	MQIACH_CHANNEL_TABLE:           "MQCHTAB",
	MQIACH_CHANNEL_TYPE:            "MQCHT",
	MQIACH_CHANNEL_TYPES:           "MQCHT",
	MQIACH_CLIENT_CHANNEL_WEIGHT:   "",
	MQIACH_CLWL_CHANNEL_PRIORITY:   "",
	MQIACH_CLWL_CHANNEL_RANK:       "",
	MQIACH_CLWL_CHANNEL_WEIGHT:     "",
	MQIACH_COMMAND_COUNT:           "",
	MQIACH_COMPRESSION_RATE:        "",
	MQIACH_COMPRESSION_TIME:        "",
	MQIACH_CONNECTION_AFFINITY:     "MQCAFTY",
	MQIACH_CURRENT_CHL:             "",
	MQIACH_CURRENT_CHL_LU62:        "",
	MQIACH_CURRENT_CHL_MAX:         "",
	MQIACH_CURRENT_CHL_TCP:         "",
	MQIACH_CURRENT_MSGS:            "",
	MQIACH_CURRENT_SEQ_NUMBER:      "",
	MQIACH_CURRENT_SHARING_CONVS:   "",
	MQIACH_DATA_CONVERSION:         "MQCDC",
	MQIACH_DATA_COUNT:              "",
	MQIACH_DEF_CHANNEL_DISP:        "MQCHLD",
	MQIACH_DEF_RECONNECT:           "MQRCN",
	MQIACH_DISC_INTERVAL:           "",
	MQIACH_DISC_RETRY:              "",
	MQIACH_DISPS_MAX:               "",
	MQIACH_DISPS_STARTED:           "",
	MQIACH_EXIT_TIME_INDICATOR:     "",
	MQIACH_HB_INTERVAL:             "",
	MQIACH_HDR_COMPRESSION:         "MQCOMPRESS",
	MQIACH_INBOUND_DISP:            "MQINBD",
	MQIACH_IN_DOUBT:                "MQIDO",
	MQIACH_IN_DOUBT_IN:             "MQIDO",
	MQIACH_IN_DOUBT_OUT:            "MQIDO",
	MQIACH_INDOUBT_STATUS:          "MQCHIDS",
	MQIACH_KEEP_ALIVE_INTERVAL:     "MQKAI",
	MQIACH_LAST_SEQ_NUMBER:         "",
	MQIACH_LISTENER_CONTROL:        "MQSVC_CONTROL",
	MQIACH_LISTENER_STATUS:         "MQSVC_STATUS",
	MQIACH_LONG_RETRIES_LEFT:       "",
	MQIACH_LONG_RETRY:              "",
	MQIACH_LONG_TIMER:              "",
	MQIACH_MATCH:                   "MQMATCH",
	MQIACH_MAX_INSTANCES:           "",
	MQIACH_MAX_INSTS_PER_CLIENT:    "",
	MQIACH_MAX_MSG_LENGTH:          "",
	MQIACH_MAX_SHARING_CONVS:       "",
	MQIACH_MAX_XMIT_SIZE:           "",
	MQIACH_MCA_STATUS:              "MQMCAS",
	MQIACH_MCA_TYPE:                "MQMCAT",
	MQIACH_MC_HB_INTERVAL:          "",
	MQIACH_MQTT_KEEP_ALIVE:         "",
	MQIACH_MR_COUNT:                "",
	MQIACH_MR_INTERVAL:             "",
	MQIACH_MSG_COMPRESSION:         "MQCOMPRESS",
	MQIACH_MSG_HISTORY:             "",
	MQIACH_MSGS:                    "",
	MQIACH_MSG_SEQUENCE_NUMBER:     "",
	MQIACH_MSGS_RCVD:               "",
	MQIACH_MSGS_SENT:               "",
	MQIACH_MULTICAST_PROPERTIES:    "MQMCP",
	MQIACH_NAME_COUNT:              "",
	MQIACH_NETWORK_PRIORITY:        "",
	MQIACH_NETWORK_TIME_INDICATOR:  "",
	MQIACH_NEW_SUBSCRIBER_HISTORY:  "MQNSH",
	MQIACH_NPM_SPEED:               "MQNPMS",
	MQIACH_PENDING_OUT:             "",
	MQIACH_PORT:                    "",
	MQIACH_PORT_NUMBER:             "",
	MQIACH_PROTOCOL:                "MQPROTO",
	MQIACH_PUT_AUTHORITY:           "MQPA",
	MQIACH_RESET_REQUESTED:         "MQCHRR",
	MQIACH_SECURITY_PROTOCOL:       "MQSECPROT",
	MQIACH_SEQUENCE_NUMBER_WRAP:    "",
	MQIACH_SESSION_COUNT:           "",
	MQIACH_SHARED_CHL_RESTART:      "MQCHSH",
	MQIACH_SHARING_CONVERSATIONS:   "",
	MQIACH_SHORT_RETRIES_LEFT:      "",
	MQIACH_SHORT_RETRY:             "",
	MQIACH_SHORT_TIMER:             "",
	MQIACH_SOCKET:                  "",
	MQIACH_SPL_PROTECTION:          "MQSPL",
	MQIACH_SSL_CLIENT_AUTH:         "MQSCA",
	MQIACH_SSL_KEY_RESETS:          "",
	MQIACH_SSL_RETURN_CODE:         "",
	MQIACH_SSLTASKS_MAX:            "",
	MQIACH_SSLTASKS_STARTED:        "",
	MQIACH_STOP_REQUESTED:          "MQCHSR",
	MQIACH_USE_CLIENT_ID:           "MQUCI",
	MQIACH_USER_SOURCE:             "MQUSRC",
	MQIACH_WARNING:                 "MQWARN",
	MQIACH_XMIT_PROTOCOL_TYPE:      "MQXPT",
	MQIACH_XMITQ_MSGS_AVAILABLE:    "",
	MQIACH_XMITQ_TIME_INDICATOR:    "MQMON",
}

func PCFAttrToPrefix(a int32) string {
	s, ok := pcfAttrMap[a]
	if ok {
		return s
	}

	return ""
}

// PCFValueToString does both steps in a single call for cases
// where it's a single integer returned. The PCFParameter
// holds int64 values, so that's what we'll use here even if we
// then have to coerce to other int types
func PCFValueToString(a int32, val int64) string {
	prefix := PCFAttrToPrefix(a)
	switch prefix {
	case "":
		return strconv.Itoa(int(val))

	default:
		// Many attributes have "special" values that do convert, but also
		// permit arbitrary integers. So first try to convert to a string.
		// If that fails, return the actual number.
		s := MQItoString(prefix, int(val))
		if s == "" {
			s = strconv.Itoa(int(val))
		}
		return s
	}
}
