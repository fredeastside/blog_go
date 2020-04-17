<template>
  <div id="header-post">
    <a id="menu-icon" href="#"
      ><font-awesome-icon :icon="['fas', 'bars']" size="lg"
    /></a>
    <a id="menu-icon-tablet" href="#"
      ><font-awesome-icon :icon="['fas', 'bars']" size="lg"
    /></a>
    <a
      id="top-icon-tablet"
      href="#"
      onclick="$('html, body').animate({ scrollTop: 0 }, 'fast');"
      style="display:none;"
      ><font-awesome-icon :icon="['fas', 'chevron-up']" size="lg"
    /></a>
    <span id="menu">
      <span id="nav">
        <ul>
          <li><router-link to="/">Home</router-link></li>
          <li><router-link to="/about">About</router-link></li>
        </ul>
      </span>
      <br />
      <span id="actions">
        <ul>
          <li>
            <a
              class="icon"
              href="#"
              onclick="$('html, body').animate({ scrollTop: 0 }, 'fast');"
              ><i
                class="fas fa-chevron-up"
                aria-hidden="true"
                onmouseover="$('#i-top').toggle();"
                onmouseout="$('#i-top').toggle();"
              ></i
            ></a>
          </li>
          <li>
            <a class="icon" href="#"
              ><i
                class="fas fa-share-alt"
                aria-hidden="true"
                onmouseover="$('#i-share').toggle();"
                onmouseout="$('#i-share').toggle();"
                onclick="$('#share').toggle();return false;"
              ></i
            ></a>
          </li>
        </ul>
        <span id="i-prev" class="info" style="display:none;"></span>
        <span id="i-next" class="info" style="display:none;"></span>
        <span id="i-top" class="info" style="display:none;"></span>
        <span id="i-share" class="info" style="display:none;"></span>
      </span>
      <br />
      <div id="share" style="display: none"></div>
    </span>
  </div>
</template>

<script>
export default {
  name: "HeaderPost"
};
</script>

<style scoped></style>
