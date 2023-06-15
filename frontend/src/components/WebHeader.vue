<template>
  <header>
    <ul id="top-title">
      <li><img id="tanzuLogo" alt="Tanzu logo" src="../assets/logo.png"></li>
      <li><h1>Tanzu Let's Get Moving</h1></li>
    </ul>
    <nav>
      <ul id="navigation-bar">
        <li class="nav-item"><router-link to="/" class="navbar-brand">Home</router-link></li>
        <li class="nav-item"><router-link to="/about" class="nav-link">About</router-link></li>
        <li class="nav-item"><router-link to="/challenges" class="nav-link">Challenges</router-link></li>
      </ul>
    </nav>
  </header>
</template>

<style scoped>
#tanzuLogo {
  width: 40%;
  height: auto;
}

#top-title li {
  display:inline;
  float: left;
}

#navigation-bar li {
  display: inline;
}

nav {
  width: 95vw;
  background-color: #0d0d0e;
  display: flex;
  align-items: center;
  justify-content: center;
}
nav .nav-item {
  color: #FFF;
  padding: 40px 70px;
  position: relative;
  text-align: center;
  border-bottom: 6px solid transparent;
  transition: 0.4s;
}
nav .nav-item.active,
nav .nav-item:hover {
  background-color: #0d0d0e;
  border-bottom-color: #1fccc0;
}
nav .nav-item a {
  color: inherit;
  text-decoration: none;
}
</style>
