# frozen_string_literal: true

module Idpd
  module V1
    class HTTP < Nonnative::HTTPClient
      def create_pipeline(params, opts = {})
        post('/pipeline', params.to_json, opts)
      end

      def get_pipeline(id, opts = {})
        get("/pipeline/#{id}", opts)
      end

      def update_pipeline(id, params, opts = {})
        put("/pipeline/#{id}", params.to_json, opts)
      end
    end
  end
end
