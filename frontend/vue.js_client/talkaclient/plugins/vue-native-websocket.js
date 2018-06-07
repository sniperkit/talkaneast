import VueNativeSock from 'vue-native-websocket'
import Vue from 'vue'
const WebSocket = require('ws').Server;

Vue.use(VueNativeSock, 'ws://localhost:2148/ws',{ 
    format: 'json',
    reconnection: true, // (Boolean) whether to reconnect automatically (false)
    reconnectionAttempts: 5, // (Number) number of reconnection attempts before giving up (Infinity),
    reconnectionDelay: 30000, // (Number) how long to initially wait before attempting a new (1000)
})