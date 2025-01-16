
export const connectToServer = () => {
    const conn = new WebSocket("ws://localhost:8080/ws")
    conn.onopen = () => {

    }
    conn.onmessage = (event) => {

    }
    sendData(conn)
}


const sendData = (conn) => {
    conn.send()
}