<template lang="pug">
  .wscn-http404-container
    .wscn-http404
      .pic-404
        img.pic-404__parent(src="@/assets/404_images/404.png" alt="404")
        img.pic-404__child.left(src="@/assets/404_images/404_cloud.png" alt="404")
        img.pic-404__child.mid(src="@/assets/404_images/404_cloud.png" alt="404")
        img.pic-404__child.right(src="@/assets/404_images/404_cloud.png" alt="404")
      .bullshit
        .bullshit__oops OOPS!
        .bullshit__info You lost
          a.link-type(href="https://www.tma.com.vn" target="_blank") TMA Solutions
        .bullshit__headline {{ message }}
        .bullshit__info Page not found, click the link below to return
        a.bullshit__return-home(href="/") Home
</template>
