# frozen_string_literal: true

When('we try to update an nonexistent pipeline') do
  pipeline = {
    pipeline: { 'name' => 'update-test', 'jobs' => [{ 'name' => 'test', 'steps' => %w[test test2] }] }
  }

  @response = Idpd::V1.server.update_pipeline(10, pipeline, opts)
end

When('we try to update an invalid pipeline') do
  pipeline = {
    pipeline: { 'jobs' => [{ 'name' => 'test', 'steps' => %w[test test2] }] }
  }

  @response = Idpd::V1.server.update_pipeline(3, pipeline, opts)
end

When('we try to update with an invalid id') do
  pipeline = {
    pipeline: { 'name' => 'update-test', 'jobs' => [{ 'name' => 'test', 'steps' => %w[test test2] }] }
  }

  @response = Idpd::V1.server.update_pipeline(0, pipeline, opts)
end

When('we update the simple pipeline') do
  pipeline = {
    pipeline: { 'name' => 'update-test', 'jobs' => [{ 'name' => 'test', 'steps' => %w[test test2] }] }
  }

  @response = Idpd::V1.server.update_pipeline(3, pipeline, opts)
end

Then('we should have an updated simple pipeline') do
  expect(@response.code).to eq(200)

  res = JSON.parse(@response.body)
  expect(res['pipeline']['name']).to eq('update-test')
end
