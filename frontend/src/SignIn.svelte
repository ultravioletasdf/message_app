<script lang="ts">
  import { WindowReload } from "../wailsjs/runtime/runtime";
  import { SignIn } from "../wailsjs/go/main/App";
  const BLANK_CHARACTER = "‎";
  let usernameOrEmail = "";
  let password = "";
  let error = BLANK_CHARACTER;
  let hasTyped = false;

  $: if (!hasTyped) {
    error = BLANK_CHARACTER;
  } else if (usernameOrEmail.length < 3) {
    error = "Username must be 3 characters or longer";
  } else if (password.length < 8) {
    error = "Password must be 8 characters or longer";
  } else {
    error = BLANK_CHARACTER;
  }
</script>

<div
  class="flex fixed top-0 left-0 z-20 justify-center items-center w-full h-full"
>
  <form
    on:submit|preventDefault
    class="flex flex-col gap-4 p-6 w-96 rounded-lg border-2 border-opacity-5 border-solid backdrop-blur-3xl border-slate-200 h-fit bg-base-300/90"
  >
    <h1 class="text-xl font-bold">Sign In</h1>
    <div class="flex flex-col gap-2">
      <label class="flex gap-2 items-center bg-opacity-50 input input-bordered">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 16 16"
          fill="currentColor"
          class="w-4 h-4 opacity-70"
          ><path
            d="M8 8a3 3 0 1 0 0-6 3 3 0 0 0 0 6ZM12.735 14c.618 0 1.093-.561.872-1.139a6.002 6.002 0 0 0-11.215 0c-.22.578.254 1.139.872 1.139h9.47Z"
          /></svg
        >
        <input
          type="text"
          class="grow"
          placeholder="Username or Email"
          bind:value={usernameOrEmail}
          on:input={() => {
            hasTyped = true;
          }}
        />
      </label>
      <label class="flex gap-2 items-center bg-opacity-50 input input-bordered">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 16 16"
          fill="currentColor"
          class="w-4 h-4 opacity-70"
          ><path
            fill-rule="evenodd"
            d="M14 6a4 4 0 0 1-4.899 3.899l-1.955 1.955a.5.5 0 0 1-.353.146H5v1.5a.5.5 0 0 1-.5.5h-2a.5.5 0 0 1-.5-.5v-2.293a.5.5 0 0 1 .146-.353l3.955-3.955A4 4 0 1 1 14 6Zm-4-2a.75.75 0 0 0 0 1.5.5.5 0 0 1 .5.5.75.75 0 0 0 1.5 0 2 2 0 0 0-2-2Z"
            clip-rule="evenodd"
          /></svg
        >
        <input
          type="password"
          class="grow"
          placeholder="Password"
          bind:value={password}
          on:input={() => {
            hasTyped = true;
          }}
        />
      </label>
      <div class="text-sm text-error">{error}</div>
    </div>
    <button
      class="btn btn-block btn-primary"
      disabled={usernameOrEmail.length < 3 || password.length < 8}
      on:click={() => {
        SignIn("a", "b");
        WindowReload();
      }}>Continue</button
    >
  </form>
</div>
<div class="fixed right-4 bottom-4 z-30 opacity-70 text-slate-400">
  Photo by
  <a
    href="https://unsplash.com/@fakurian?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash"
    class="hover:text-white">Milad Fakurian</a
  >
  on
  <a
    href="https://unsplash.com/photos/a-blue-abstract-sculpture-against-a-purple-background-eoUIrRBB61I?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash"
    class="hover:text-white">Unsplash</a
  >
</div>
<div
  class="fixed top-0 left-0 z-10 bg-top blur-md scale-105 w-dvw h-dvh"
  style="background-image: url('https://images.unsplash.com/photo-1707246933257-6789cca3e5f8?q=80&w=1932&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D')"
></div>
