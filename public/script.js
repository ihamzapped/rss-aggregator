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
  try {
    const res = await fetch(`${apiUrl}users`);
    const users = await res.json();

    const el = document.getElementById("users");
    el.textContent = JSON.stringify(users, null, 2);
  } catch (error) {
    console.log(error);
  }
}

async function login(e) {
  e.preventDefault();

  headers.Authorization = authPrefix + document.getElementById("apiKey").value;

  fetchFeed();
}

async function fetchFeed() {
  try {
    const res = await fetch(`${apiUrl}user-feeds`, {
      method: "GET",
      headers,
    });

    if (!res.ok) await throwResErr(res);

    const postsArray = await res.json();
    console.log(postsArray);

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
        publishedAt: post?.published_at,
      });

      feed.innerHTML += html;
    }
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

// UTILITIES

function renderTemplate(template, data) {
  return template.replace(/{{\s*([a-zA-Z0-9_]+)\s*}}/g, function (match, p1) {
    return data[p1] || match;
  });
}

async function throwResErr(res) {
  const err = {
    status: res.status,
    statusText: res.statusText,
    ...(await res.json()),
  };

  alertObj(err);
  throw err;
}

function alertObj(obj) {
  alert(JSON.stringify(obj, null, 2));
}
