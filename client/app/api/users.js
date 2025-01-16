export const getUsers = async () => {
  const usersList = document.querySelector(".users-list");
  try {
    let res = await fetch("/api/info");
    let data = await res.json();
    data.Clients.forEach((user) => {
      const div = document.createElement("div");
      div.classList.add("user-item");
      div.id = user.userId;
      div.setAttribute("onclick", `target(${user.userId})`)
      div.innerHTML = `${user.nickname} <span class="${
        user.status ? "online" : "offline"
      }"></span>`;
      usersList.append(div)
    });
  } catch (err) {
    console.error(err);
  }
};

