import { getRequest, prepandMessage, createUser } from "../utils/helpers.js";
import { messages } from "../services/users.js";
export let myNickname;
export let myId;
export const limit = 10;

export const getUsers = async () => {
  try {
    let res = await fetch("/api/info");
    let data = await res.json();
    myNickname = data.Nickname;
    myId = data.UserId;
    data.Clients.forEach((user) => {
      createUser(user);
    });
  } catch (err) {
    console.error(err);
  }
};

export const loadMessages = async (container, id) => {
  const url =
    `/api/client/${id}?` + new URLSearchParams({ offset: messages[id] });
  const data = await getRequest(url);
  if (data.length != 0) {
    data.forEach((msg) => {
      prepandMessage(
        msg.content,
        msg.sender,
        new Date(msg.creation).toLocaleTimeString()
      );
    });
    container.scroll(0, 85 * limit);
  }
};
