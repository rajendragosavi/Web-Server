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

helm push my-go-app-chart-0.1.0.tgz oci://<your-registry>/<your-repo>/charts  // Ifg you want to push Helm Chart to oci repo

helm install my-go-app-release oci://<your-registry>/<your-repo>/charts/my-go-app-chart --version 0.1.0 --set image.repository=<your-dockerhub-username>/my-go-app --set image.tag=latest  // Pull chart from oci repo.

helm install my-go-app-release my-gh-pages-repo/my-go-app-chart --set image.repository=<your-dockerhub-username>/my-go-app --set image.tag=latest

```






## How to use chart

```
helm repo add my-git-repo https://raw.githubusercontent.com/rajendragosavi/Web-Server/master/my-go-app-chart
```


```
helm repo update
```

```
helm install my-go-app-release my-git-repo/my-go-app-chart \
  --set image.repository=<your-dockerhub-username>/my-go-app \
  --set image.tag=latest
```



```
docker build -t my-go-app:latest .
```





![PlantUML model](bLLHRzem47xFhx1u0gcexpnKocZQib7QaR9fZoQv1K_6Zco7ZVtwdHr8OYVGhX9Iu_DzzztTT-UIAiCoLb9Xo4Q3YihEmK8K2hHEH5vA9end-ootQ0Pjc1HVAwZ0iywbWZJRGLxn07CbSrX3mRHHnvuXBa0OHEu3DjQGH6gegNHoGDEdIeVMPofe4J1mJiYlXt--7Db2JE-2x6SahRPxPkcDVydjFHZAkBuXwJAUH3Gpx40Dz4eMV0-WVgESvTOLWVlNUsMUaTTgovdUZGNyHkiaok5YzCZhmo3QN0fTxM5SAYOLC-mlv0wrmVHR6G1skpaZ2AcEyG7hGhSSh4sFCrhIBUE835WiBcdUmVR4qFsOJk-yI3Eo1bCfuJmMugTSuhLkvXvIo3A0YikI7ydfW8c2nBomkU_s0_LoKfcHJMbxrEZDAc2Htv1HdAQBnrLoFxeXL6jM2CWTnMwMZU3E__IpY-TVkjaq7WFbzcQaLKChoU-OoPb31zFlOzH03QAULT6Y1bshPt4812afmD42alo1hB9I8uauC2N5_h8FbbMvqKHlcQ66k_TwO-wbxVgmtgnMYzKZjg5KqWs1LKdO16zLz4wc7YFRal_ZC7zvVbqcRoW6yh8ceuTuiNJ2G3OLD_eMb99gyf6uWR3hOaVjFXkaq-bPRQerwzVWlEDwMDTKZfkcFfXq9Tn2rAGZokiEFg1U88y8_Ws7SNi2Ft6crJ4r2cCMnphOvEms5EMq0poteJ0jBJVfxZvPoe9bxhRbQS-OtLBkP03ZqjPj19UaywVaVhDCHfDB6E_iZdNojlvXHUi1N2glDGFNLX4VeFzEw9CsFnvOv97z-JkgUJ23cMvVX5V0PoLd6UKSlrA4U08Sz2rGCnPkHzLlGZ2irzw5wwFIlW6T5_qwgsSf6APkho0vRdRZWCavpG4YJQ5QxxhIG-zJ5yTZ5kWrHMBsMQeLUoMFkhlsxDHUckUt2lRo0CHMDrpQrqhgjVnIKVy1)
