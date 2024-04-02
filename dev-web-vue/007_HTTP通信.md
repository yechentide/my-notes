# HTTP通信

## FetchAPI

```javascript
// データ受信
const getData = async () => {
  try {
    const response = await fetch("http://........");
    if (!response.ok) throw new Error('Fetch Error 404 etc.');
    const data = await response.json();
    // ...
  } catch (error) {
    console.error(error);
  }
}

// データ送信
const sendData = async (data) => {
  try {
    const options = {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
          "Content-Type": "application/json"
      }
    };
    const response = await fetch("http://......", options);
    if (!response.ok) throw new Error('Fetch Error 404 etc.');
    console.log("Done!");
  } catch (error) {
    console.error(error);
  }
}
```
