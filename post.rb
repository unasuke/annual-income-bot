require 'logger'
require 'twitter'
logger = Logger.new(STDOUT)

config = {
  consumer_key:        ENV['CONSUMER_KEY'],
  consumer_secret:     ENV['CONSUMER_SECRET'],
  access_token:        ENV['ACCESS_TOKEN'],
  access_token_secret: ENV['ACCESS_TOKEN_SECRET'],
}

rest_client = Twitter::REST::Client.new(config)
logger.info 'REST client initialized'

streaming_client = Twitter::Streaming::Client.new(config)
logger.info 'Streaming client initialized'

streaming_client.user do |tweet|
  if tweet.is_a?(Twitter::DirectMessage) && %r[\A年収\z].match?(tweet.text)
    logger.info "Recieved DM #{tweet.text} from #{tweet.sender.screen_name}"
    rest_client.create_direct_message(tweet.sender, "#{ENV['ANNUAL_INCOME']}万円")
  end
end
