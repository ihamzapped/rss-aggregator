const apiUrl = "http://localhost:8000/v1/";

// To construct Authorization header
// i.e "authPrefix + " " + apiKey"
const authPrefix = "ApiKey ";

const headers = {
  "Content-Type": "application/json",
  Authorization: "",
};

(function INIT() {
  fetchUsers();
})();

async function fetchUsers() {
  const getUsers = await (await fetch(`${apiUrl}users`)).json();
  const el = document.getElementById("users");

  el.textContent = JSON.stringify(getUsers, null, 2);
}

async function login(e) {
  e.preventDefault();

  headers.Authorization = authPrefix + document.getElementById("apiKey").value;

  try {
    const res = await fetch(`${apiUrl}user-feeds`, {
      method: "GET",
      headers,
    });
    const result = await res.json();
    console.log(result);
    setFeed(result);
  } catch (error) {
    console.log(error);
  }
}

async function createUser(e) {
  e.preventDefault();

  const formData = {
    name: document.getElementById("username").value,
  };

  try {
    await fetch(`${apiUrl}user`, {
      method: "POST",
      headers,
      body: JSON.stringify(formData),
    });

    fetchUsers();
  } catch (error) {
    console.log(error);
  }
}

function setFeed(postsArray) {
  const loginScreen = document.getElementById("login-screen");
  const mainScreen = document.getElementById("main-screen");
  const feed = document.getElementById("feed");
  const template = document.getElementById("post-template").innerHTML;
  loginScreen.classList.add("hidden");
  mainScreen.classList.remove("hidden");

  for (const post of postsArray) {
    const html = renderTemplate(template, {
      title: post?.title,
      description: post?.description,
      url: post?.url,
      updatedAt: post?.updated_at,
      createdAt: post?.created_at,
    });

    feed.innerHTML += html;
  }
}

function renderTemplate(template, data) {
  return template.replace(/{{\s*([a-zA-Z0-9_]+)\s*}}/g, function (match, p1) {
    return data[p1] || match;
  });
}
