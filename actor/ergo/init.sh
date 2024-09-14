# An actor-based Framework with network transparency
# for creating event-driven architecture in Golang.
# Inspired by Erlang. Zero dependencies.

ergo -init MyNode \
      -with-app MyApp \
      -with-sup MyApp:MySup \
      -with-actor MySup:MyActor \
      -with-web MyWeb \
      -with-actor MyActor2 \
      -with-observer


