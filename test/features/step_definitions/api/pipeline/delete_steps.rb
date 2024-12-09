# frozen_string_literal: true

When('we delete the simple pipeline') do
  @id = JSON.parse(@response.body)['pipeline']['id'].to_i
  @response = Idpd.http.delete_pipeline(@id, Idpd.options)
end

When('we try to delete an nonexistent pipeline') do
  @response = Idpd.http.delete_pipeline(10, Idpd.options)
end

When('we try to delete with an invalid id') do
  @response = Idpd.http.delete_pipeline(0, Idpd.options)
end

Then('we should have deleted simple pipeline') do
  expect(@response.code).to eq(200)

  res = JSON.parse(@response.body)
  expect(res['pipeline']['id']).to eq(@id)
end
