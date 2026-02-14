> [!IMPORTANT]
> **Render Cold Starts:** Services may be sleeping.
> Before testing, please **wake up the services** by activating the ping in **Better Stack** or manually hitting the `/health` endpoints:
> * `GET https://nexus-auth-service.onrender.com/health`
> * `GET https://nexus-catalog.onrender.com/health`
>
> Please allow **~50-60 seconds** for the initial response.
