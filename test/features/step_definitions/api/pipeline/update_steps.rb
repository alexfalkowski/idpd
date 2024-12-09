# frozen_string_literal: true

When('we try to update an nonexistent pipeline') do
  pipeline = {
    pipeline: { 'name' => 'update-test', 'jobs' => [{ 'name' => 'test', 'steps' => %w[test test2] }] }
  }

  @response = Idpd.http.update_pipeline(10, pipeline.to_json, Idpd.options)
end

When('we try to update an invalid pipeline') do
  pipeline = {
    pipeline: { 'jobs' => [{ 'name' => 'test', 'steps' => %w[test test2] }] }
  }

  @response = Idpd.http.update_pipeline(3, pipeline.to_json, Idpd.options)
end

When('we try to update with an invalid id') do
  pipeline = {
    pipeline: { 'name' => 'update-test', 'jobs' => [{ 'name' => 'test', 'steps' => %w[test test2] }] }
  }

  @response = Idpd.http.update_pipeline(0, pipeline.to_json, Idpd.options)
end

When('we update the simple pipeline') do
  pipeline = {
    pipeline: { 'name' => 'update-test', 'jobs' => [{ 'name' => 'test', 'steps' => %w[test test2] }] }
  }

  @id = JSON.parse(@response.body)['pipeline']['id'].to_i
  @response = Idpd.http.update_pipeline(@id, pipeline.to_json, Idpd.options)
end

When('we try update a pipeline with a bad payload') do
  @id = JSON.parse(@response.body)['pipeline']['id'].to_i
  @response = Idpd.http.update_pipeline(@id, Base64.encode64('test'), Idpd.options)
end

Then('we should have an updated simple pipeline') do
  expect(@response.code).to eq(200)

  res = JSON.parse(@response.body)
  pipeline = res['pipeline']

  expect(pipeline['name']).to eq('update-test')
  expect(pipeline['id']).to eq(@id)
end
