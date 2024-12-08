# frozen_string_literal: true

module Idpd
  class HTTP < Nonnative::HTTPClient
    def create_pipeline(pipeline, opts = {})
      post('/pipeline', pipeline, opts)
    end

    def get_pipeline(id, opts = {})
      get("/pipeline/#{id}", opts)
    end

    def update_pipeline(id, pipeline, opts = {})
      put("/pipeline/#{id}", pipeline, opts)
    end

    def delete_pipeline(id, opts = {})
      delete("/pipeline/#{id}", opts)
    end

    def trigger_pipeline(id, opts = {})
      post("/pipeline/#{id}/trigger", '', opts)
    end
  end
end
