import os
import logging
from mobile_app_react_native import config

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s] %(message)s",
    datefmt="%Y-%m-%d %H:%M:%S"
)

class Application:
    def __init__(self):
        self.config = config.Config()

    def start(self):
        if not os.path.exists(self.config.app_data_path):
            os.makedirs(self.config.app_data_path)

        # Initialize services
        self.services = {
            "database": self.init_database(),
            "network": self.init_network(),
            "storage": self.init_storage()
        }

        # Start main loop
        while True:
            self.loop()

    def loop(self):
        # Handle events
        for service_name, service in self.services.items():
            service.handle_events()

        # Update services
        for service_name, service in self.services.items():
            service.update()

    def init_database(self):
        from mobile_app_react_native import database
        return database.Database(self.config.database_path)

    def init_network(self):
        from mobile_app_react_native import network
        return network.Network(self.config.network_url)

    def init_storage(self):
        from mobile_app_react_native import storage
        return storage.Storage(self.config.storage_path)

if __name__ == "__main__":
    app = Application()
    app.start()