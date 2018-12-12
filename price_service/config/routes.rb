Rails.application.routes.draw do
  namespace :api do
    namespace :v1 do
      resources :prices, only: %i[] do
        collection do
          post :calculate
        end
      end
    end
  end
end
