# frozen_string_literal: true

require 'securerandom'
require 'yaml'
require 'base64'

require 'idpd/http'

module Idpd
  class << self
    def config
      @config ||= Nonnative.configurations('.config/server.yml')
    end

    def token
      Nonnative::Header.auth_bearer(Base64.decode64(File.read('secrets/token')))
    end

    def http
      @http ||= Idpd::HTTP.new(Nonnative.configuration.url)
    end

    def options
      {
        headers: {
          request_id: SecureRandom.uuid, user_agent: 'IDP-ruby-client/1.0 HTTP/1.0',
          content_type: :json, accept: :json
        }.merge(Idpd.token),
        read_timeout: 10, open_timeout: 10
      }
    end
  end
end
