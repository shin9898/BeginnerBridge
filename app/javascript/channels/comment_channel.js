import consumer from "channels/consumer"

if(location.pathname.match(/\/posts\/\d/)){

  consumer.subscriptions.create("CommentChannel", {
    connected() {
      // Called when the subscription is ready for use on the server
    },

    disconnected() {
      // Called when the subscription has been terminated by the server
    },

    received(data) {
      const comments = document.getElementById("comments")
      comments.insertAdjacentHTML('beforeend', data.html);
      const commentForm = document.getElementById("comment-form")
      commentForm.reset();
    }
  });
}
