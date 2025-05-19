from fastapi import FastAPI, WebSocket, WebSocketDisconnect
from fastapi.middleware.cors import CORSMiddleware
from PIL import Image
import io
import base64
import numpy as np
import os
from dotenv import load_dotenv
from datetime import datetime
from ai.face_recognizer import FaceRecognizer
from utils.database import Database

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


load_dotenv()  # .env dosyasını yükle

db_config = {
    "host": os.getenv("DB_HOST"),
    "user": os.getenv("DB_USER"),
    "password": os.getenv("DB_PASSWORD"),
    "database": os.getenv("DB_NAME")
}


db = Database(**db_config)
face_recognizer = FaceRecognizer(db_config)

@app.websocket("/ws/attendance")
async def websocket_endpoint(websocket: WebSocket):
    await websocket.accept()
    print("WebSocket connection established.")
    try:
        while True:
            try:
                data = await websocket.receive_text()
                print("Received data (first 100 chars):", data[:100])

                header, encoded = data.split(",", 1)
                image_data = base64.b64decode(encoded)
                image = Image.open(io.BytesIO(image_data)).convert("RGB")
                rgb_image = np.array(image)

                os.makedirs("data/attendance_images", exist_ok=True)
                filename = f"data/attendance_images/{datetime.now().strftime('%Y%m%d_%H%M%S')}.jpg"
                image.save(filename)
                print(f"Image saved as {filename}")

                bgr_image = rgb_image[:, :, ::-1]

                student = face_recognizer.recognize_faces(bgr_image)
                if student:
                    
                    success = db.create_attendance_if_needed_for_student(student.id)
                    if success:
                        print("[INFO] Attendance recorded.")
                    else:
                        print("[WARNING] Attendance was already created or an error occurred.")
                    
                    response = {
                        "status": 200,
                        "data": student.to_dict()
                    }
                else:
                    response = {
                        "status": 404,
                        "message": "Face not recognized or student not found."
                    }

                await websocket.send_json(response)
                print("Response sent:", response)

            except Exception as e:
                print("Error during image handling or recognition:", e)
                await websocket.send_json({
                    "status": 500,
                    "message": "Internal server error during image processing."
                })

    except WebSocketDisconnect:
        print("WebSocket connection closed.")
