export let client = null

export const WS = {

    connect: function() {
        this.client = new WebSocket('ws://localhost:8080/api/ws/connect');
    },

    close: function() {
        this.client.close()
    },
}