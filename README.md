# Web-Server
WebRTC application

1. Create Helm Chart templates

``` helm create my-go-app-chart ```

2. Update Chart.yaml file based on your requirments.


3. Build Helm Package. This will create a .tgz archive of your chart 
cd Web-Server/my-go-app-chart

``` helm package . ```

4. Create index.yaml | This will create an index.yaml file in the same directory as the .tgz archive.


``` helm repo index . ```

## Host on GitHub Pages (Recommended for Sharing)

GitHub Pages is a free and easy way to host your Helm chart.

* Create a docs Directory:
    * Create a docs directory in the root of your repository.
    * Move the .tgz archive and index.yaml file into the docs directory.

* Enable GitHub Pages:

    * Go to your repository on GitHub.
    * Go to "Settings" -> "Pages".
    * Under "Source," select "Deploy from a branch."
    * Select the branch you want to use (e.g., master or main) and the /docs folder.
    * Click "Save."
    * GitHub Pages will build and deploy your site.
    * Use the GitHub Pages URL:
    * Once GitHub Pages is deployed, you'll get a URL (e.g., https://rajendragosavi.github.io/Web-Server/).

Use this URL with helm repo add:

``` 
helm repo add my-gh-pages-repo https://rajendragosavi.github.io/Web-Server/
helm repo update
helm push my-go-app-chart-0.1.0.tgz oci://<your-registry>/<your-repo>/charts  // Push Helm Chart to oci repo
helm install my-go-app-release oci://<your-registry>/<your-repo>/charts/my-go-app-chart --version 0.1.0 --set image.repository=<your-dockerhub-username>/my-go-app --set image.tag=latest  // Pull chart from oci repo.

helm install my-go-app-release my-gh-pages-repo/my-go-app-chart --set image.repository=<your-dockerhub-username>/my-go-app --set image.tag=latest
```






## How to use chart

helm repo add my-git-repo https://raw.githubusercontent.com/rajendragosavi/Web-Server/master/my-go-app-chart


helm repo update


helm install my-go-app-release my-git-repo/my-go-app-chart \
  --set image.repository=<your-dockerhub-username>/my-go-app \
  --set image.tag=latest




```
docker build -t my-go-app:latest .
```