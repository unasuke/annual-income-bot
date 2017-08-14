FROM ruby:2.4.1-slim

WORKDIR /app

RUN apt update && apt install --assume-yes g++ make

COPY Gemfile Gemfile.lock /app/
RUN bundle install

COPY post.rb /app
CMD bundle exec ruby post.rb
